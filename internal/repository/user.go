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
	err = r.db.Model(&user).Association("Company").Find(&user.Company)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user.
func (r *userRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}
