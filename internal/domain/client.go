package domain

import (
	"gorm.io/gorm"
)

// Client represents a client entity associated with a company
type Client struct {
	gorm.Model
	CompanyID      uint                `gorm:"not null"` // Foreign key linking to the Company model
	Company        Company             `gorm:"foreignKey:CompanyID"`
	CompanyName    string              `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	Representative string              `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	Phone          string              `gorm:"type:varchar(20);not null" validate:"required,e164"`
	PostalCode     string              `gorm:"type:varchar(10);not null" validate:"required,len=10"`
	Address        string              `gorm:"type:varchar(255);not null" validate:"required,max=255"`
	BankAccounts   []ClientBankAccount `gorm:"foreignKey:ClientID"` // One-to-many relationship
	Invoices       []InvoiceData       `gorm:"foreignKey:ClientID"`
}
