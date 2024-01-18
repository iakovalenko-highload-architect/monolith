package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"monolith/cmd"
	"monolith/internal/api/handlers/login"
	"monolith/internal/api/handlers/user_get_by_id"
	"monolith/internal/api/handlers/user_register"
	"monolith/internal/storage"
	"monolith/internal/usecase/auth_manager"
	"monolith/internal/usecase/hash_manager"
	"monolith/internal/usecase/token_manager"
	"monolith/internal/usecase/user_manager"
	"monolith/internal/utils/jwt"
)

func main() {
	userStorage := storage.New(cmd.MustInitPostgresql())

	hashManager := hash_manager.New(cmd.MustInitHasherConfig())
	tokenManager := token_manager.New(jwt.New(), cmd.MustInitTokenManagerConfig())
	authManager := auth_manager.New(userStorage, hashManager, tokenManager)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	loginHandler := login.New(authManager)
	userRegisterHandler := user_register.New(authManager)
	userGetByIDHandler := user_get_by_id.New(user_manager.New(userStorage))

	e.POST("/login", loginHandler.Handle)
	e.POST("/user/register", userRegisterHandler.Handle)
	e.GET("/user/get/:id", userGetByIDHandler.Handle)

	e.Logger.Fatal(e.Start(":8080"))
}
