package domain

import (
	"gorm.io/gorm"
)

// User represents a user entity in the system
type User struct {
	gorm.Model
	Username  string  `gorm:"unique;not null" validate:"required"`
	Email     string  `gorm:"unique;not null" validate:"required,email"`
	Password  string  `gorm:"type:varchar(255);not null" validate:"required,min=8"`
	CompanyID uint    // Foreign key for Company
	Company   Company `gorm:"foreignKey:CompanyID"`
}
