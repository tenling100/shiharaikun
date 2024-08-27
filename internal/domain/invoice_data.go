package domain

import (
	"time"

	"gorm.io/gorm"
)

// InvoiceData represents an invoice linked to a company and a client
type InvoiceData struct {
	gorm.Model
	CompanyID      uint      `gorm:"not null"` // Foreign key linking to the Company model
	Company        Company   `gorm:"foreignKey:CompanyID"`
	ClientID       uint      `gorm:"not null"` // Foreign key linking to the Client model
	Client         Client    `gorm:"foreignKey:ClientID"`
	IssueDate      time.Time `gorm:"not null"`                    // Date when the invoice was issued
	PaymentAmount  float64   `gorm:"type:decimal(10,2);not null"` // Total payment amount
	Fee            float64   `gorm:"type:decimal(10,2)"`          // Fee associated with the invoice
	FeeRate        float64   `gorm:"type:decimal(5,2)"`           // Fee rate (percentage)
	Tax            float64   `gorm:"type:decimal(10,2)"`          // Tax amount
	TaxRate        float64   `gorm:"type:decimal(5,2)"`           // Tax rate (percentage)
	InvoiceAmount  float64   `gorm:"type:decimal(10,2);not null"` // Total invoice amount
	PaymentDueDate time.Time `gorm:"not null"`                    // Due date for payment
	Status         string    `gorm:"type:varchar(20);not null"`   // Status of the invoice
}
