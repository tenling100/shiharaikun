package repository

import "github.com/tenling100/shiharaikun/internal/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
}

type CompanyRepository interface {
	CreateCompany(company *domain.Company) error
	GetCompanyByID(id uint) (*domain.Company, error)
	UpdateCompany(company *domain.Company) error
}

type InvoiceDataRepository interface {
	CreateInvoiceData(invoiceData *domain.InvoiceData) error
	GetInvoiceDataByID(id uint) (*domain.InvoiceData, error)
	UpdateInvoiceData(invoiceData *domain.InvoiceData) error
}
