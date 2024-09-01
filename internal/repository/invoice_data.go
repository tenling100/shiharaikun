package repository

import (
	"github.com/tenling100/shiharaikun/internal/domain"
	"gorm.io/gorm"
)

type invoiceDataRepository struct {
	db *gorm.DB
}

func NewInvoiceDataRepository(db *gorm.DB) *invoiceDataRepository {
	return &invoiceDataRepository{db: db}
}

// CreateInvoiceData creates a new invoice data.
func (r *invoiceDataRepository) CreateInvoiceData(invoiceData *domain.InvoiceData) error {
	return r.db.Create(invoiceData).Error
}

// GetInvoiceDataByID retrieves an invoice data by its ID.
func (r *invoiceDataRepository) GetInvoiceDataByID(id uint) (*domain.InvoiceData, error) {
	var invoiceData domain.InvoiceData
	err := r.db.Where("id = ?", id).First(&invoiceData).Error
	if err != nil {
		return nil, err
	}
	return &invoiceData, nil
}

// UpdateInvoiceData updates an invoice data.
func (r *invoiceDataRepository) UpdateInvoiceData(invoiceData *domain.InvoiceData) error {
	return r.db.Save(invoiceData).Error
}

// GetInvoicesByDateRange retrieves invoices by date range.
func (r *invoiceDataRepository) GetInvoicesByDateRange(startDate, endDate string) ([]domain.InvoiceData, error) {
	var invoices []domain.InvoiceData
	err := r.db.Where("issue_date BETWEEN ? AND ?", startDate, endDate).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return invoices, nil
}
