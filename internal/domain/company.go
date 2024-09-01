package domain

import (
	"time"

	_ "gorm.io/gorm"
)

type Company struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Representative string    `gorm:"type:varchar(100);not null"`
	Phone          string    `gorm:"type:varchar(20);not null"`
	PostalCode     string    `gorm:"type:varchar(20);not null"`
	Address        string    `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
