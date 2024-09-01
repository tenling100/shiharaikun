package usecase

import (
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/repository"
)

type invoiceUseCaseImpl struct {
	companyRepository     repository.CompanyRepository
	invoiceDataRepository repository.InvoiceDataRepository
}

func NewInvoiceUseCaseImpl(
	companyRepository repository.CompanyRepository,
	invoiceDataRepository repository.InvoiceDataRepository,
) InvoiceUseCase {
	return &invoiceUseCaseImpl{
		companyRepository:     companyRepository,
		invoiceDataRepository: invoiceDataRepository,
	}
}

// CreateInvoice creates a new invoice.
func (u *invoiceUseCaseImpl) CreateInvoice(invoice *domain.InvoiceData) error {
	err := u.invoiceDataRepository.CreateInvoiceData(invoice)
	if err != nil {
		return err
	}
	return nil
}

// GetInvoiceByUserID retrieves an invoice by its user ID.
