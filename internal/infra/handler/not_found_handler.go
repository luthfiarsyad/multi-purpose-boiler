package handler

import (
	"base-app/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This handler is handling not found path, so everytime the requested path is not available, it will directing here.
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the error
		logger.LogError(c.Request.Context(), logger.LogEvent{
			HTTPStatus: http.StatusNotFound,
			Message:    "Resource not found",
			Data:       gin.H{"path": c.Request.URL.Path},
		}, nil)

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Resource not found",
			"data":    nil,
		})
	}
}
