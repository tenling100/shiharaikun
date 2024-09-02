package domain

import (
	"time"

	pb "github.com/tenling100/shiharaikun/api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	_ "gorm.io/gorm"
)

// InvoiceData represents an invoice linked to a company and a client
type InvoiceData struct {
	ID             uint      `gorm:"primaryKey"`                                     // Unique identifier
	CompanyID      uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key linking to the Company model
	Company        *Company  `gorm:"foreignKey:CompanyID"`                           // Company associated with the invoice
	ClientID       uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key linking to the Client model
	Client         *Company  `gorm:"foreignKey:ClientID"`                            //
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

// ConvertToInvoiceData converts an InvoiceData object to a pb.Invoice object
func ConvertToInvoiceData(invoice *InvoiceData) *pb.Invoice {
	return &pb.Invoice{
		Id: wrapperspb.UInt64(uint64(invoice.ID)),
		Company: &pb.Company{
			Id:             wrapperspb.UInt64(uint64(invoice.CompanyID)),
			Name:           invoice.Company.Name,
			Representative: invoice.Company.Representative,
			Phone:          invoice.Company.Phone,
			PostalCode:     invoice.Company.PostalCode,
			Address:        invoice.Company.Address,
		},
		Client: &pb.Company{
			Id:             wrapperspb.UInt64(uint64(invoice.ClientID)),
			Name:           invoice.Client.Name,
			Representative: invoice.Client.Representative,
			Phone:          invoice.Client.Phone,
			PostalCode:     invoice.Client.PostalCode,
			Address:        invoice.Client.Address,
		},
		IssueDate:      timestamppb.New(invoice.IssueDate),
		PaymentAmount:  float32(invoice.PaymentAmount),
		Fee:            float32(invoice.Fee),
		FeeRate:        float32(invoice.FeeRate),
		Tax:            float32(invoice.Tax),
		TaxRate:        float32(invoice.TaxRate),
		InvoiceAmount:  float32(invoice.InvoiceAmount),
		PaymentDueDate: timestamppb.New(invoice.PaymentDueDate),
		Status:         pb.InvoiceStatus(pb.InvoiceStatus_value[invoice.Status]),
		CreatedAt:      timestamppb.New(invoice.CreatedAt),
		UpdatedAt:      timestamppb.New(invoice.UpdatedAt),
	}

}
