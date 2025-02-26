package handler

import (
	"base-app/internal/domain"
	"base-app/internal/service"
	"base-app/pkg/logger"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// Use Gin's request context and add required values
	ctx := c.Request.Context()

	// Set transaction ID if not present
	transactionID, ok := ctx.Value(logger.TransactionIDKey).(string)
	if !ok {
		transactionID = uuid.New().String()
		ctx = context.WithValue(ctx, logger.TransactionIDKey, transactionID)
	}

	// Set start time for process time calculation
	startTime := time.Now()
	ctx = context.WithValue(ctx, logger.StartTimeKey, startTime)

	logger.LogInfo(ctx, logger.LogEvent{
		HTTPStatus: http.StatusOK,
		Message:    "Received CreateUser request",
	})

	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.LogError(ctx, logger.LogEvent{
			HTTPStatus: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       req,
		}, err)
		c.JSON(http.StatusBadRequest, domain.Response{
			Status: http.StatusBadRequest,
			Data:   gin.H{"error": err.Error()},
		})
		return
	}

	user := domain.CreateUserRequest{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.userService.CreateUser(ctx, &user); err != nil {
		logger.LogError(ctx, logger.LogEvent{
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Failed to create user",
			Data:       req,
		}, err)
		c.JSON(http.StatusInternalServerError, domain.Response{
			Status: http.StatusInternalServerError,
			Data:   gin.H{"error": err.Error()},
		})
		return
	}

	logger.LogInfo(ctx, logger.LogEvent{
		HTTPStatus: http.StatusCreated,
		Message:    "User created successfully",
		Data:       user,
	})

	c.JSON(http.StatusCreated, domain.Response{
		Status: http.StatusCreated,
		Data:   gin.H{"message": "User created", "user": user},
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// Use Gin's request context
	ctx := c.Request.Context()

	// Set transaction ID if not present
	transactionID, ok := ctx.Value(logger.TransactionIDKey).(string)
	if !ok {
		transactionID = uuid.New().String()
		ctx = context.WithValue(ctx, logger.TransactionIDKey, transactionID)
	}

	// Set start time for process time calculation
	startTime := time.Now()
	ctx = context.WithValue(ctx, logger.StartTimeKey, startTime)

	logger.LogInfo(ctx, logger.LogEvent{
		HTTPStatus: http.StatusOK,
		Message:    "Received GetUser request",
	})

	id := c.Param("id")

	userId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logger.LogError(ctx, logger.LogEvent{
			HTTPStatus: http.StatusBadRequest,
			Message:    "Invalid user ID",
			Data:       gin.H{"id": id},
		}, err)
		c.JSON(http.StatusBadRequest, domain.Response{
			Status: http.StatusBadRequest,
			Data:   gin.H{"error": "Invalid user ID"},
		})
		return
	}

	user, err := h.userService.GetUserByID(ctx, uint(userId))
	if err != nil {
		logger.LogError(ctx, logger.LogEvent{
			HTTPStatus: http.StatusInternalServerError,
			Message:    "Failed to fetch user",
			Data:       gin.H{"id": userId},
		}, err)
		c.JSON(http.StatusInternalServerError, domain.Response{
			Status: http.StatusInternalServerError,
			Data:   gin.H{"error": err.Error()},
		})
		return
	}

	logger.LogInfo(ctx, logger.LogEvent{
		HTTPStatus: http.StatusOK,
		Message:    "User retrieved successfully",
		Data:       user,
	})

	c.JSON(http.StatusOK, domain.Response{
		Status: http.StatusOK,
		Data:   user,
	})
}
