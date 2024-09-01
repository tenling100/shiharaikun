package usecase

import "github.com/tenling100/shiharaikun/internal/domain"

type InvoiceUseCase interface {
	CreateInvoice(invoice *domain.InvoiceData) error
}

type CompanyUsecase interface {
	CreateCompany(company *domain.Company) error
}

type UserUsecase interface {
	CreateUser(user *domain.User) error
}
