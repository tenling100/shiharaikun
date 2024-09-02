package usecase

import (
	"fmt"
	"math"

	"github.com/tenling100/shiharaikun/internal/config"
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/repository"
)

type invoiceUsecaseImpl struct {
	invoiceDataRepository repository.InvoiceDataRepository
	companyRepository     repository.CompanyRepository
	userRepository        repository.UserRepository
	env                   *config.Env
}

func NewInvoiceUsecaseImpl(
	invoiceDataRepository repository.InvoiceDataRepository,
	companyRepository repository.CompanyRepository,
	userRepository repository.UserRepository,
	env *config.Env,
) InvoiceUseCase {
	return &invoiceUsecaseImpl{
		invoiceDataRepository: invoiceDataRepository,
		companyRepository:     companyRepository,
		userRepository:        userRepository,
		env:                   env,
	}
}

// CreateInvoice creates a new invoice.
func (u *invoiceUsecaseImpl) CreateInvoice(invoice *domain.InvoiceData) error {
	// calculate fee and tax
	invoice.FeeRate = u.env.FeeRate
	invoice.TaxRate = u.env.TaxRate
	invoice.Fee = (invoice.InvoiceAmount * invoice.FeeRate)
	invoice.Tax = (invoice.InvoiceAmount * invoice.TaxRate)
	invoice.PaymentAmount = math.Ceil(invoice.InvoiceAmount + invoice.Fee + invoice.Tax)
	company, err := u.companyRepository.GetCompanyByID(invoice.CompanyID)
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("company not found, please create company before creating invoice, %v", err)
		}
		return err
	}
	client, err := u.companyRepository.GetCompanyByID(invoice.ClientID)
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("client not found, please create client before creating invoice, %v", err)
		}
		return err
	}
	invoice.Company = company
	invoice.Client = client

	err = u.invoiceDataRepository.CreateInvoiceData(invoice)
	if err != nil {
		return err
	}

	return nil
}

func (u *invoiceUsecaseImpl) GetInvoicesByDateRange(startDate, endDate string) ([]domain.InvoiceData, error) {
	invoices, err := u.invoiceDataRepository.GetInvoicesByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}
	if len(invoices) == 0 {
		return nil, fmt.Errorf("no invoices found")
	}

	for i := range invoices {
		company, err := u.companyRepository.GetCompanyByID(invoices[i].CompanyID)
		if err != nil {
			return nil, err
		}
		client, err := u.companyRepository.GetCompanyByID(invoices[i].ClientID)
		if err != nil {
			return nil, err
		}
		invoices[i].Company = company
		invoices[i].Client = client
	}
	return invoices, nil
}
