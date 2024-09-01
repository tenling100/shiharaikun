package usecase

import "github.com/tenling100/shiharaikun/internal/domain"

type InvoiceUseCase interface {
	CreateInvoice(invoice *domain.Invoice) error
}
