package domain

import (
	_ "gorm.io/gorm"
)

// ClientBankAccount represents a bank account associated with a client
type ClientBankAccount struct {
	ID            uint   `gorm:"primaryKey"`
	CompanyID     uint   `gorm:"not null;foreignKey:CompanyID"` // Foreign key linking to the Client model
	BankName      string `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	BranchName    string `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	AccountNumber string `gorm:"type:varchar(20);not null" validate:"required,max=20"`
	AccountName   string `gorm:"type:varchar(100);not null" validate:"required,max=100"`
}
