package routes

import (
	"base-app/internal/infra/handler"
	"base-app/internal/infra/repo"
	"base-app/internal/service"
	"base-app/pkg/db"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	NotFoundHandler gin.HandlerFunc
	*handler.UserHandler
}

func InitWiring() Handlers {
	// Init DB
	Db := db.InitDB()

	// Initialize Repository
	userRepo := repo.NewUserRepo(Db)

	// Initialize Service
	userService := service.NewUserService(userRepo)

	// Initialize Handler
	userHandler := handler.NewUserHandler(userService)
	noRouteHandler := handler.NotFoundHandler()

	return Handlers{
		UserHandler:     userHandler,
		NotFoundHandler: noRouteHandler,
	}
}
