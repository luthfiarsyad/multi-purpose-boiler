package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouting() *gin.Engine {
	// Init Routing
	r := gin.New()

	h := InitWiring()

	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUser)

	// return the routes
	return r
}
