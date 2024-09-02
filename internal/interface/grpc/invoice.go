package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/usecase"
)

type InvoiceServer struct {
	invoiceUsecase usecase.InvoiceUseCase
	companyUseCase usecase.CompanyUsecase
	userUseCase    usecase.UserUsecase
	pb.UnimplementedInvoiceServiceServer
}

func NewInvoiceServer(
	invoiceUsecase usecase.InvoiceUseCase,
	companyUseCase usecase.CompanyUsecase,
	userUseCase usecase.UserUsecase,
) *InvoiceServer {
	return &InvoiceServer{
		invoiceUsecase: invoiceUsecase,
		companyUseCase: companyUseCase,
		userUseCase:    userUseCase,
	}
}

func (s *InvoiceServer) CreateInvoice(ctx context.Context, in *pb.InvoiceRequest) (*pb.InvoiceResponse, error) {
	if in.IssueDate.AsTime().After(time.Now()) {
		return nil, fmt.Errorf("issue date must be in the past")
	}
	if in.PaymentDueDate.AsTime().Before(time.Now()) {
		return nil, fmt.Errorf("payment due date must be in the future")
	}
	if in.InvoiceAmount <= 0 {
		return nil, fmt.Errorf("invoice amount must be greater than 0")
	}
	invoice := &domain.InvoiceData{
		CompanyID:      uint(in.CompanyId),
		ClientID:       uint(in.ClientId),
		IssueDate:      in.IssueDate.AsTime(),
		InvoiceAmount:  float64(in.InvoiceAmount),
		Status:         in.Status.String(),
		PaymentDueDate: in.PaymentDueDate.AsTime(),
	}

	err := s.invoiceUsecase.CreateInvoice(invoice)
	if err != nil {
		return nil, err
	}

	return &pb.InvoiceResponse{
		Invoice: domain.ConvertToInvoiceData(invoice),
	}, nil
}

func (s *InvoiceServer) GetInvoicesByDateRange(ctx context.Context, in *pb.DateRangeRequest) (*pb.InvoicesResponse, error) {
	if in.StartDate.AsTime().After(time.Now()) || in.EndDate.AsTime().After(time.Now()) {
		return nil, fmt.Errorf("start date and end date must be in the past")
	}
	if in.StartDate.AsTime().After(in.EndDate.AsTime()) {
		return nil, fmt.Errorf("start date must be before end date")
	}
	invoices, err := s.invoiceUsecase.GetInvoicesByDateRange(in.StartDate.AsTime().String(), in.EndDate.AsTime().String())
	if err != nil {
		return nil, err
	}
	var pbInvoices []*pb.Invoice
	for _, invoice := range invoices {
		pbInvoice := domain.ConvertToInvoiceData(&invoice)
		pbInvoices = append(pbInvoices, pbInvoice)
	}
	return &pb.InvoicesResponse{
		Invoices: pbInvoices,
	}, nil
}

func (s *InvoiceServer) CreateCompany(ctx context.Context, in *pb.Company) (*pb.Company, error) {
	company := &domain.Company{
		Name:           in.Name,
		Representative: in.Representative,
		Phone:          in.Phone,
		PostalCode:     in.PostalCode,
		Address:        in.Address,
	}
	err := s.companyUseCase.CreateCompany(company)
	if err != nil {
		return nil, err
	}
	return &pb.Company{
		Id:             wrapperspb.UInt64(uint64(company.ID)),
		Name:           company.Name,
		Representative: company.Representative,
		Phone:          company.Phone,
		PostalCode:     company.PostalCode,
		Address:        company.Address,
	}, nil
}
func (s *InvoiceServer) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	user := &domain.User{
		Username:  in.Name,
		Email:     in.Email,
		CompanyID: uint(in.CompanyId),
	}
	err := s.userUseCase.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        wrapperspb.UInt64(uint64(user.ID)),
		Name:      user.Username,
		Email:     user.Email,
		CompanyId: uint64(user.CompanyID),
	}, nil
}
