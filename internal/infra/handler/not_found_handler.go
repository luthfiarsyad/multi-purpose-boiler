package handler

import (
	"base-app/pkg/logger"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// This handler is handling not found path, so everytime the requested path is not available, it will directing here.
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, logger.StartTimeKey, startTime)
		// Log the error
		logger.LogError(ctx, logger.LogEvent{
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
