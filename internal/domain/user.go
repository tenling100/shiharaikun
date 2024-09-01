package domain

import (
	"time"

	_ "gorm.io/gorm"
)

// User represents a user entity in the system
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null" validate:"required"`
	Email     string    `gorm:"unique;not null" validate:"required,email"`
	Password  string    `gorm:"type:varchar(255);not null"validate:"required,min=8"`
	CompanyID uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Company   *Company  `gorm:"foreignKey:CompanyID;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
