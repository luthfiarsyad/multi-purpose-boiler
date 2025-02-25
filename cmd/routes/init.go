package routes

import (
	"base-app/internal/infra/handler"
	"base-app/internal/infra/repo"
	"base-app/internal/service"
	"base-app/pkg/db"
)

type Handlers struct {
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

	return Handlers{UserHandler: userHandler}
}
