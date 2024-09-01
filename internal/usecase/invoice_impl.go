package usecase

import (
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/repository"
)

type invoiceUsecaseImpl struct {
	companyRepository     repository.CompanyRepository
	invoiceDataRepository repository.InvoiceDataRepository
}

func NewInvoiceUsecaseImpl(
	companyRepository repository.CompanyRepository,
	invoiceDataRepository repository.InvoiceDataRepository,
) InvoiceUseCase {
	return &invoiceUsecaseImpl{
		companyRepository:     companyRepository,
		invoiceDataRepository: invoiceDataRepository,
	}
}

// CreateInvoice creates a new invoice.
func (u *invoiceUsecaseImpl) CreateInvoice(invoice *domain.InvoiceData) error {
	err := u.invoiceDataRepository.CreateInvoiceData(invoice)
	if err != nil {
		return err
	}
	return nil
}
