//go:generate mockgen -destination=mock/mock.go -source=interface.go
package usecase

import "github.com/tenling100/shiharaikun/internal/domain"

type InvoiceUseCase interface {
	CreateInvoice(invoice *domain.InvoiceData) error
	GetInvoicesByDateRange(startDate, endDate string) ([]domain.InvoiceData, error)
}

type CompanyUsecase interface {
	CreateCompany(company *domain.Company) error
}

type UserUsecase interface {
	CreateUser(user *domain.User) error
}
