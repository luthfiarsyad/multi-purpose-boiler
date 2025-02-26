package repo

import (
	"base-app/internal/domain"
	"base-app/pkg/logger"
	"context"

	"gorm.io/gorm"
)

// UserRepo interface defines the contract for user-related database operations.
type UserRepo interface {
	CreateUser(ctx context.Context, user *domain.CreateUserRequest) error
	GetUserByID(ctx context.Context, id uint) (*domain.GetUserResponse, error)
}

// userRepo implements the UserRepo interface.
type userRepo struct {
	DB *gorm.DB
}

// NewUserRepo creates a new UserRepo instance.
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db}
}

// CreateUser inserts a new user into the database using raw SQL.
func (r *userRepo) CreateUser(ctx context.Context, user *domain.CreateUserRequest) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	err := r.DB.WithContext(ctx).Exec(query, user.Name, user.Email).Error

	if err != nil {
		logger.LogError(ctx, logger.LogEvent{
			HTTPStatus: 500,
			Message:    "Failed to create user",
			Data:       user,
                        LogPoint:   "database-response",
		}, err)
		return err
	}

	logger.LogInfo(ctx, logger.LogEvent{
		HTTPStatus: 201,
		Message:    "User created successfully",
		Data:       user,
		LogPoint:   "database-response",
	})
	return nil
}

// GetUserByID fetches a user by ID using raw SQL.
func (r *userRepo) GetUserByID(ctx context.Context, id uint) (*domain.GetUserResponse, error) {
	var user domain.GetUserResponse
	query := `SELECT id, name, email FROM users WHERE id = ?`
	err := r.DB.WithContext(ctx).Raw(query, id).Scan(&user).Error

	if err != nil {
		logger.LogError(ctx, logger.LogEvent{
			HTTPStatus: 500,
			Message:    "Failed to fetch user",
			Data:       id,
                        LogPoint:   "database-response",
		}, err)
		return nil, err
	}

	logger.LogInfo(ctx, logger.LogEvent{
		HTTPStatus: 200,
		Message:    "User fetched successfully",
		Data:       user,
                LogPoint:   "database-response",
	})
	return &user, nil
}
