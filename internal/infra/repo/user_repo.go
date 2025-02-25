package repo

import (
	"base-app/internal/domain"

	"gorm.io/gorm"
)

// UserRepo interface defines the contract for user-related database operations.
type UserRepo interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	// Add other repository methods as needed (e.g., UpdateUser, DeleteUser, GetAllUsers)
}

// userRepo implements the UserRepo interface.  It's a struct that holds the DB connection.
type userRepo struct {
	DB *gorm.DB
}

// NewUserRepo creates a new UserRepo instance.  This is the factory function.
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{DB: db} // Return the interface type
}

func (r *userRepo) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepo) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
