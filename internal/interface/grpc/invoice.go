package grpc

import (
	"context"

	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/usecase"
)

type InvoiceServer struct {
	invoiceUsecase usecase.InvoiceUseCase
	pb.UnimplementedInvoiceServiceServer
}

func NewInvoiceServer(invoiceUsecase usecase.InvoiceUseCase) *InvoiceServer {
	return &InvoiceServer{
		invoiceUsecase: invoiceUsecase,
	}
}

func (s *InvoiceServer) CreateInvoice(ctx context.Context, in *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	return nil, nil
}

func (s *InvoiceServer) GetInvoicesByDateRange(ctx context.Context, in *pb.DateRangeRequest) (*pb.InvoicesResponse, error) {
	return nil, nil
}
