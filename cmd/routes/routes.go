package routes

import (
	_ "base-app/docs" // Import the generated Swagger docs

	"github.com/gin-gonic/gin"
	ginFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouting() *gin.Engine {
	// Init Routing
	r := gin.New()

	h := InitWiring()

	// Swagger docs route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(ginFiles.Handler))

	// App Routes
	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUser)

	// return the routes
	return r
}
