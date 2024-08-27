package domain

import (
	"gorm.io/gorm"
)

// Company represents a company entity
type Company struct {
	gorm.Model
	Name           string        `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	Representative string        `gorm:"type:varchar(100);not null" validate:"required,max=100"`
	Phone          string        `gorm:"type:varchar(20);not null" validate:"required,e164"`
	PostalCode     string        `gorm:"type:varchar(10);not null" validate:"required,len=10"`
	Address        string        `gorm:"type:varchar(255);not null" validate:"required,max=255"`
	Clients        []Client      `gorm:"foreignKey:CompanyID"` // One-to-many relationship
	Invoices       []InvoiceData `gorm:"foreignKey:CompanyID"`
}
