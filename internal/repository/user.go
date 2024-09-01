package repository

import (
	"gorm.io/gorm"

	"github.com/tenling100/shiharaikun/internal/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// CreateUser creates a new user.
func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

// GetUserByEmail retrieves a user by its email.
func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user.
func (r *userRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}

// GetUserByID retrieves a user by its ID.
func (r *userRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
