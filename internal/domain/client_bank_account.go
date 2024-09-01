package domain

import (
	"time"

	_ "gorm.io/gorm"
)

// ClientBankAccount represents a bank account associated with a client
type ClientBankAccount struct {
	ID            uint      `gorm:"primaryKey"`
	CompanyID     uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Company       Company   `gorm:"foreignKey:CompanyID;"` // Correctly defined foreign key
	BankName      string    `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	BranchName    string    `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	AccountNumber string    `gorm:"type:varchar(20);not null" validate:"required,max=20"`
	AccountName   string    `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
