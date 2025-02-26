package service

import (
	"base-app/internal/domain"
	"base-app/internal/infra/repo"
	"context"
)

// UserService interface defines the user-related service operations.
type UserService interface {
	CreateUser(ctx context.Context, user *domain.CreateUserRequest) error
	GetUserByID(ctx context.Context, id uint) (*domain.GetUserResponse, error)
}

type userService struct {
	userRepo repo.UserRepo
}

// NewUserService creates a new UserService, injecting the UserRepo.
func NewUserService(ur repo.UserRepo) UserService {
	return &userService{userRepo: ur}
}

func (s *userService) CreateUser(ctx context.Context, user *domain.CreateUserRequest) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*domain.GetUserResponse, error) {
	return s.userRepo.GetUserByID(ctx, id)
}
