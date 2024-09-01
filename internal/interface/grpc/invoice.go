package grpc

import (
	"context"
	"time"

	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/usecase"
)

type invoiceServer struct {
	invoiceUsecase usecase.InvoiceUseCase
}

func NewInvoiceServer(invoiceUsecase *usecase.InvoiceUseCase) *invoiceServer {
	return &invoiceServer{
		invoiceUsecase: *invoiceUsecase,
	}
}

func (s *invoiceServer) CreateInvoice(ctx context.Context, in *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	invoice := &domain.Invoice{
		Status:    in.invoice.Status,
		Amount:    in.invoice.Amount,
		RepayDate: in.invoice.RepayDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.invoiceUsecase.CreateInvoice(in.Invoice)
	return nil, nil
}

func (s *invoiceServer) GetInvoicesByDateRange(ctx context.Context, in *pb.DateRangeRequest) (*pb.InvoicesResponse, error) {
	return nil, nil
}
