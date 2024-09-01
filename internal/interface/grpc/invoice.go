package grpc

import (
	"context"
	"fmt"

	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	invoice := &domain.Invoice{
		Status:    in.Status.String(),
		Amount:    uint64(in.Amount),
		RepayDate: in.RepaymentDate.AsTime(),
	}
	err := s.invoiceUsecase.CreateInvoice(invoice)
	if err != nil {
		return nil, err
	}
	return &pb.InvoiceResponse{
		Invoice: &pb.Invoice{
			Id:            fmt.Sprint(invoice.ID),
			Status:        pb.InvoiceStatus(pb.InvoiceStatus_value[invoice.Status]),
			Amount:        uint64(invoice.Amount),
			RepaymentDate: timestamppb.New(invoice.RepayDate),
		},
	}, nil
}

func (s *InvoiceServer) GetInvoicesByDateRange(ctx context.Context, in *pb.DateRangeRequest) (*pb.InvoicesResponse, error) {
	return nil, nil
}
