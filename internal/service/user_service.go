package service

import (
	"base-app/internal/domain"
	"base-app/internal/infra/repo"
)

// UserService interface defines the user-related service operations.
type UserService interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
}

type userService struct {
	userRepo repo.UserRepo // Dependency Injection: Inject the UserRepo interface.
}

// NewUserService creates a new UserService, injecting the UserRepo.
func NewUserService(ur repo.UserRepo) UserService { //Notice the interface type here
	return &userService{userRepo: ur}
}

func (s *userService) CreateUser(user *domain.User) error {
	//Service layer logic if needed, e.g. validation, business rules
	return s.userRepo.CreateUser(user) // Call the repository's CreateUser method
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	//Service layer logic if needed, e.g. authorization, caching
	return s.userRepo.GetUserByID(id) // Call the repository's GetUserByID method
}
