package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
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
		Invoice: &pb.Invoice{
			Id: wrapperspb.UInt64(uint64(invoice.ID)),
			Company: &pb.Company{
				Id:             wrapperspb.UInt64(uint64(invoice.CompanyID)),
				Name:           invoice.Company.Name,
				Representative: invoice.Company.Representative,
				Phone:          invoice.Company.Phone,
				PostalCode:     invoice.Company.PostalCode,
				Address:        invoice.Company.Address,
			},
			Client: &pb.Company{
				Id:         wrapperspb.UInt64(uint64(invoice.ClientID)),
				Name:       invoice.Client.Name,
				Phone:      invoice.Client.Phone,
				PostalCode: invoice.Client.PostalCode,
				Address:    invoice.Client.Address,
			},
			IssueDate:      timestamppb.New(invoice.IssueDate),
			PaymentAmount:  float32(invoice.PaymentAmount),
			Fee:            float32(invoice.Fee),
			FeeRate:        float32(invoice.FeeRate),
			Tax:            float32(invoice.Tax),
			TaxRate:        float32(invoice.TaxRate),
			InvoiceAmount:  float32(invoice.InvoiceAmount),
			PaymentDueDate: timestamppb.New(invoice.PaymentDueDate),
			Status:         pb.InvoiceStatus(pb.InvoiceStatus_value[invoice.Status]),
		},
	}, nil
}

func (s *InvoiceServer) GetInvoicesByDateRange(ctx context.Context, in *pb.DateRangeRequest) (*pb.InvoicesResponse, error) {
	return nil, nil
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
