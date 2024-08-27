package grpc

import (
	"context"

	pb "github.com/tenling100/shiharaikun/api"
)

type invoiceServer struct {
}

func NewInvoiceServer() *invoiceServer {
	return &invoiceServer{}
}

func (s *invoiceServer) CreateInvoice(ctx context.Context, in *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	return nil, nil
}

func (s *invoiceServer) GetInvoicesByDateRange(ctx context.Context, in *pb.DateRangeRequest) (*pb.InvoicesResponse, error) {
	return nil, nil
}
