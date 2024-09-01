package usecase

import (
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
	invoice.Fee = math.Round(invoice.InvoiceAmount * invoice.FeeRate)
	invoice.Tax = math.Round(invoice.InvoiceAmount * invoice.TaxRate)
	invoice.PaymentAmount = invoice.InvoiceAmount + invoice.Fee + invoice.Tax
	company, err := u.companyRepository.GetCompanyByID(invoice.CompanyID)
	if err != nil {
		return err
	}
	client, err := u.companyRepository.GetCompanyByID(invoice.ClientID)
	if err != nil {
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
