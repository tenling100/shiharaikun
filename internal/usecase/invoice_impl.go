package usecase

import (
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/repository"
)

type InvoiceUseCaseImpl struct {
	InvoiceRepository repository.InvoiceRepository
}

func NewInvoiceUseCaseImpl(invoiceRepository repository.InvoiceRepository) *InvoiceUseCaseImpl {
	return &InvoiceUseCaseImpl{InvoiceRepository: invoiceRepository}
}

// CreateInvoice creates a new invoice.
func (u *InvoiceUseCaseImpl) CreateInvoice(invoice *domain.Invoice) error {

	return u.InvoiceRepository.CreateInvoice(invoice)
}
