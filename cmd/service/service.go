package main

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/wagslane/go-rabbitmq"

	"monolith/cmd"
	"monolith/internal/api/handlers/dialog_user_id_list"
	"monolith/internal/api/handlers/dialog_user_id_send"
	"monolith/internal/api/handlers/friend_delete_by_id"
	"monolith/internal/api/handlers/friend_set_by_id"
	"monolith/internal/api/handlers/login"
	"monolith/internal/api/handlers/post_create"
	"monolith/internal/api/handlers/post_delete"
	"monolith/internal/api/handlers/post_feed"
	"monolith/internal/api/handlers/post_get_by_id"
	"monolith/internal/api/handlers/post_update"
	"monolith/internal/api/handlers/user_get_by_id"
	"monolith/internal/api/handlers/user_register"
	"monolith/internal/api/handlers/user_search"
	"monolith/internal/cache"
	"monolith/internal/clients/dialog"
	customMiddlewares "monolith/internal/middlewares"
	db "monolith/internal/storage"
	"monolith/internal/usecase/auth_manager"
	"monolith/internal/usecase/dialog_manager"
	"monolith/internal/usecase/friend_manager"
	"monolith/internal/usecase/hash_manager"
	"monolith/internal/usecase/post_manager"
	"monolith/internal/usecase/token_manager"
	"monolith/internal/usecase/user_manager"
	"monolith/internal/utils/jwt"
)

func main() {
	ctx := context.Background()

	storage := db.New(cmd.MustInitPostgresql())
	storageRO := db.New(cmd.MustInitPostgresqlRO())

	feedCache := cache.New(cmd.MustInitRedis())

	rabbit := cmd.MustInitRabbit()
	defer rabbit.Close()

	publisher, err := rabbitmq.NewPublisher(
		rabbit,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("post-created"),
		rabbitmq.WithPublisherOptionsExchangeKind("topic"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		panic(err)
	}
	defer publisher.Close()

	dialogClient := dialog.New(os.Getenv("SERVICE_DIALOG_URL"))

	hashManager := hash_manager.New(cmd.MustInitHasherConfig())
	tokenManager := token_manager.New(jwt.New(), cmd.MustInitTokenManagerConfig())
	authManager := auth_manager.New(storage, hashManager, tokenManager)
	dialogManager := dialog_manager.New(storage, dialogClient)
	friendManager := friend_manager.New(storage)
	postManager := post_manager.New(storage, feedCache, friendManager, publisher)
	if err := postManager.InitFeedCache(ctx); err != nil {
		panic(err)
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	loginHandler := login.New(authManager)
	userRegisterHandler := user_register.New(authManager)
	userGetByIDHandler := user_get_by_id.New(user_manager.New(storageRO))
	userSearchHandler := user_search.New(user_manager.New(storageRO))
	dialogUserIdSendHandler := dialog_user_id_send.New(dialogManager)
	dialogUserIdListHandler := dialog_user_id_list.New(dialogManager)
	friendSetByIDHandler := friend_set_by_id.New(friendManager)
	friendDeleteByIDHandler := friend_delete_by_id.New(friendManager)
	postCreateHandler := post_create.New(postManager)
	postUpdateHandler := post_update.New(postManager)
	postDeleteHandler := post_delete.New(postManager)
	postGetHandler := post_get_by_id.New(postManager)
	feedGetHandler := post_feed.New(feedCache)

	e.POST("/login", loginHandler.Handle)
	e.POST("/user/register", userRegisterHandler.Handle)
	e.GET("/user/get/:id", userGetByIDHandler.Handle)
	e.GET("/user/search", userSearchHandler.Handle)
	e.GET("/post/get/:id", postGetHandler.Handle)

	middlewares := customMiddlewares.New(tokenManager)

	dialogGroup := e.Group("/dialog")
	dialogGroup.Use(middlewares.Auth)
	dialogGroup.POST("/:user_id/send", dialogUserIdSendHandler.Handle)
	dialogGroup.GET("/:user_id/list", dialogUserIdListHandler.Handle)

	friendGroup := e.Group("/friend")
	friendGroup.Use(middlewares.Auth)
	friendGroup.PUT("/set/:user_id", friendSetByIDHandler.Handle)
	friendGroup.PUT("/delete/:user_id", friendDeleteByIDHandler.Handle)

	postGroup := e.Group("/post")
	postGroup.Use(middlewares.Auth)
	postGroup.POST("/create", postCreateHandler.Handle)
	postGroup.PUT("/update", postUpdateHandler.Handle)
	postGroup.PUT("/delete/:id", postDeleteHandler.Handle)
	postGroup.GET("/feed", feedGetHandler.Handle)

	e.Logger.Fatal(e.Start(":8080"))
}
