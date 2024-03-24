package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"monolith/cmd"
	"monolith/internal/api/handlers/dialog_user_id_list"
	"monolith/internal/api/handlers/dialog_user_id_send"
	"monolith/internal/api/handlers/login"
	"monolith/internal/api/handlers/user_get_by_id"
	"monolith/internal/api/handlers/user_register"
	"monolith/internal/api/handlers/user_search"
	"monolith/internal/clients/dialog"
	customMiddlewares "monolith/internal/middlewares"
	"monolith/internal/storage"
	"monolith/internal/usecase/auth_manager"
	"monolith/internal/usecase/dialog_manager"
	"monolith/internal/usecase/hash_manager"
	"monolith/internal/usecase/token_manager"
	"monolith/internal/usecase/user_manager"
	"monolith/internal/utils/jwt"
)

func main() {
	userStorage := storage.New(cmd.MustInitPostgresql())
	userStorageRO := storage.New(cmd.MustInitPostgresqlRO())

	dialogClient := dialog.New(os.Getenv("SERVICE_DIALOG_URL"))

	hashManager := hash_manager.New(cmd.MustInitHasherConfig())
	tokenManager := token_manager.New(jwt.New(), cmd.MustInitTokenManagerConfig())
	authManager := auth_manager.New(userStorage, hashManager, tokenManager)
	dialogManager := dialog_manager.New(userStorage, dialogClient)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	loginHandler := login.New(authManager)
	userRegisterHandler := user_register.New(authManager)
	userGetByIDHandler := user_get_by_id.New(user_manager.New(userStorageRO))
	userSearchHandler := user_search.New(user_manager.New(userStorageRO))
	dialogUserIdSendHandler := dialog_user_id_send.New(dialogManager)
	dialogUserIdListHandler := dialog_user_id_list.New(dialogManager)

	e.POST("/login", loginHandler.Handle)
	e.POST("/user/register", userRegisterHandler.Handle)
	e.GET("/user/get/:id", userGetByIDHandler.Handle)
	e.GET("/user/search", userSearchHandler.Handle)

	middlewares := customMiddlewares.New(tokenManager)

	dialogGroup := e.Group("/dialog")
	dialogGroup.Use(middlewares.Auth)
	dialogGroup.POST("/:user_id/send", dialogUserIdSendHandler.Handle)
	dialogGroup.GET("/:user_id/list", dialogUserIdListHandler.Handle)

	e.Logger.Fatal(e.Start(":8080"))
}
