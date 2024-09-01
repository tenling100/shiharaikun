package domain

import (
	"time"

	_ "gorm.io/gorm"
)

// InvoiceData represents an invoice linked to a company and a client
type InvoiceData struct {
	ID             uint      `gorm:"primaryKey"`                                     // Unique identifier
	CompanyID      uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key linking to the Company model
	Company        Company   `gorm:"foreignKey:CompanyID"`                           // Company associated with the invoice
	ClientID       uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key linking to the Client model
	Client         Company   `gorm:"foreignKey:ClientID"`                            //
	IssueDate      time.Time `gorm:"not null"`                                       // Date when the invoice was issued
	PaymentAmount  float64   `gorm:"type:decimal(10,2);not null"`                    // Total payment amount
	Fee            float64   `gorm:"type:decimal(10,2)"`                             // Fee associated with the invoice
	FeeRate        float64   `gorm:"type:decimal(5,2)"`                              // Fee rate (percentage)
	Tax            float64   `gorm:"type:decimal(10,2)"`                             // Tax amount
	TaxRate        float64   `gorm:"type:decimal(5,2)"`                              // Tax rate (percentage)
	InvoiceAmount  float64   `gorm:"type:decimal(10,2);not null"`                    // Total invoice amount
	PaymentDueDate time.Time `gorm:"not null"`                                       // Due date for payment
	Status         string    `gorm:"type:varchar(20);not null"`                      // Status of the invoice
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
