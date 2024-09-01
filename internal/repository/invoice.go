package repository

import (
	"gorm.io/gorm"

	"github.com/tenling100/shiharaikun/internal/domain"
)

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *invoiceRepository {
	return &invoiceRepository{db: db}
}

// CreateInvoice creates a new invoice.
func (r *invoiceRepository) CreateInvoice(invoice *domain.Invoice) error {

	return r.db.Create(invoice).Error
}

// GetInvoiceByID retrieves an invoice by its ID.
func (r *invoiceRepository) GetInvoiceByID(id uint) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := r.db.Where("id = ?", id).First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

// UpdateInvoice updates an invoice.
func (r *invoiceRepository) UpdateInvoice(invoice *domain.Invoice) error {
	return r.db.Save(invoice).Error
}
