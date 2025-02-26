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

	// Serve OpenAPI YAML as a static file in /docs/
	r.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")

	// Serve Swagger UI and point it to the OpenAPI 3.0 YAML file
	r.GET("/swagger/*any", ginSwagger.WrapHandler(ginFiles.Handler, ginSwagger.URL("/docs/openapi.yaml")))

	// App Routes
	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUser)

	// return the routes
	return r
}
