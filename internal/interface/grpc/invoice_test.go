package grpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/domain"
	mock_usecase "github.com/tenling100/shiharaikun/internal/usecase/mock"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestInvoiceServer_CreateInvoice(t *testing.T) {
	now := time.Now()
	ctrl := gomock.NewController(t)
	tests := []struct {
		name    string
		request *pb.InvoiceRequest
		want    *pb.InvoiceResponse
		wantErr bool
	}{
		{
			name: "normal case",
			request: &pb.InvoiceRequest{
				CompanyId:      1,
				ClientId:       2,
				IssueDate:      timestamppb.New(now.Truncate(time.Hour * 24 * 3)),
				InvoiceAmount:  1000,
				Status:         pb.InvoiceStatus_UNPAID,
				PaymentDueDate: timestamppb.New(now.Add(time.Hour * 24 * 5)),
			},
			want: &pb.InvoiceResponse{
				Invoice: &pb.Invoice{
					Id:             wrapperspb.UInt64(1),
					IssueDate:      timestamppb.New(now.Truncate(time.Hour * 24 * 3)),
					InvoiceAmount:  1000,
					Status:         pb.InvoiceStatus_UNPAID,
					PaymentDueDate: timestamppb.New(now.Add(time.Hour * 24 * 5)),
					Company: &pb.Company{
						Id:             wrapperspb.UInt64(1),
						Name:           "test",
						Representative: "test",
						Phone:          "test",
						PostalCode:     "test",
						Address:        "test",
					},
					Client: &pb.Company{
						Id:             wrapperspb.UInt64(2),
						Name:           "test",
						Representative: "test",
						Phone:          "test",
						PostalCode:     "test",
						Address:        "test",
					},
					CreatedAt: timestamppb.New(now),
					UpdatedAt: timestamppb.New(now),
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			request: &pb.InvoiceRequest{
				CompanyId:      1,
				ClientId:       2,
				IssueDate:      timestamppb.New(now.Truncate(time.Hour * 24 * 3)),
				InvoiceAmount:  1000,
				Status:         pb.InvoiceStatus_UNPAID,
				PaymentDueDate: timestamppb.New(now.Add(time.Hour * 24 * 5)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock_invoiceUsecase := mock_usecase.NewMockInvoiceUseCase(ctrl)
			mock_companyUseCase := mock_usecase.NewMockCompanyUsecase(ctrl)
			mock_userUseCase := mock_usecase.NewMockUserUsecase(ctrl)

			invoice := &domain.InvoiceData{
				CompanyID:      uint(tt.request.CompanyId),
				ClientID:       uint(tt.request.ClientId),
				IssueDate:      tt.request.IssueDate.AsTime(),
				InvoiceAmount:  float64(tt.request.InvoiceAmount),
				Status:         tt.request.Status.String(),
				PaymentDueDate: tt.request.PaymentDueDate.AsTime(),
				Company:        nil,
				Client:         nil,
			}
			if tt.wantErr {
				mock_invoiceUsecase.EXPECT().CreateInvoice(invoice).Return(fmt.Errorf("error"))
			} else {
				mock_invoiceUsecase.EXPECT().CreateInvoice(invoice).
					Do(func(invoice *domain.InvoiceData) {
						invoice.ID = 1
						invoice.Company = &domain.Company{
							ID:             1,
							Name:           "test",
							Representative: "test",
							Phone:          "test",
							PostalCode:     "test",
							Address:        "test",
						}
						invoice.Client = &domain.Company{
							ID:             2,
							Name:           "test",
							Representative: "test",
							Phone:          "test",
							PostalCode:     "test",
							Address:        "test",
						}
						invoice.CreatedAt = now
						invoice.UpdatedAt = now
					}).
					Return(nil)
			}

			s := &InvoiceServer{
				invoiceUsecase: mock_invoiceUsecase,
				companyUseCase: mock_companyUseCase,
				userUseCase:    mock_userUseCase,
			}
			got, err := s.CreateInvoice(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvoiceServer.CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, protocmp.Transform()); diff != "" {
				t.Errorf("InvoiceServer.CreateInvoice() diff = %v", diff)
			}
		})
	}
}

func TestInvoiceServer_GetInvoicesByDateRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name    string
		request *pb.DateRangeRequest
		want    *pb.InvoicesResponse
		wantErr bool
	}{
		{
			name: "normal case",
			request: &pb.DateRangeRequest{
				StartDate: timestamppb.New(time.Now().Truncate(time.Hour * 24 * 10)),
				EndDate:   timestamppb.New(time.Now().Truncate(time.Hour * 24 * 5)),
			},
			want: &pb.InvoicesResponse{
				Invoices: []*pb.Invoice{
					{
						Id:             wrapperspb.UInt64(1),
						IssueDate:      timestamppb.New(time.Now()),
						InvoiceAmount:  1000,
						Status:         pb.InvoiceStatus_UNPAID,
						PaymentDueDate: timestamppb.New(time.Now()),
						Company: &pb.Company{
							Id:             wrapperspb.UInt64(1),
							Name:           "test",
							Representative: "test",
							Phone:          "test",
							PostalCode:     "test",
							Address:        "test",
						},
						Client: &pb.Company{
							Id:             wrapperspb.UInt64(2),
							Name:           "test",
							Representative: "test",
							Phone:          "test",
							PostalCode:     "test",
							Address:        "test",
						},
						CreatedAt: timestamppb.New(time.Now()),
						UpdatedAt: timestamppb.New(time.Now()),
					},
					{
						Id:             wrapperspb.UInt64(2),
						IssueDate:      timestamppb.New(time.Now()),
						InvoiceAmount:  1000,
						Status:         pb.InvoiceStatus_UNPAID,
						PaymentDueDate: timestamppb.New(time.Now()),
						Company: &pb.Company{
							Id:             wrapperspb.UInt64(1),
							Name:           "test",
							Representative: "test",
							Phone:          "test",
							PostalCode:     "test",
							Address:        "test",
						},
						Client: &pb.Company{
							Id:             wrapperspb.UInt64(2),
							Name:           "test",
							Representative: "test",
							Phone:          "test",
							PostalCode:     "test",
							Address:        "test",
						},
						CreatedAt: timestamppb.New(time.Now()),
						UpdatedAt: timestamppb.New(time.Now()),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			request: &pb.DateRangeRequest{
				StartDate: timestamppb.New(time.Now().Truncate(time.Hour * 24 * 10)),
				EndDate:   timestamppb.New(time.Now().Truncate(time.Hour * 24 * 5)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock_invoiceUsecase := mock_usecase.NewMockInvoiceUseCase(ctrl)
			mock_companyUseCase := mock_usecase.NewMockCompanyUsecase(ctrl)
			mock_userUseCase := mock_usecase.NewMockUserUsecase(ctrl)

			if tt.wantErr {
				mock_invoiceUsecase.EXPECT().GetInvoicesByDateRange(tt.request.StartDate.AsTime().String(), tt.request.EndDate.AsTime().String()).Return(nil, fmt.Errorf("error"))
			} else {
				invoices := []domain.InvoiceData{}
				for _, invoicepb := range tt.want.Invoices {
					invoice := domain.InvoiceData{
						ID:             uint(invoicepb.Id.Value),
						CompanyID:      uint(invoicepb.Company.Id.Value),
						ClientID:       uint(invoicepb.Client.Id.Value),
						IssueDate:      invoicepb.IssueDate.AsTime(),
						InvoiceAmount:  float64(invoicepb.InvoiceAmount),
						Status:         invoicepb.Status.String(),
						PaymentDueDate: invoicepb.PaymentDueDate.AsTime(),
						Company: &domain.Company{
							ID:             uint(invoicepb.Company.Id.Value),
							Name:           invoicepb.Company.Name,
							Representative: invoicepb.Company.Representative,
							Phone:          invoicepb.Company.Phone,
							PostalCode:     invoicepb.Company.PostalCode,
							Address:        invoicepb.Company.Address,
						},
						Client: &domain.Company{
							ID:             uint(invoicepb.Client.Id.Value),
							Name:           invoicepb.Client.Name,
							Representative: invoicepb.Client.Representative,
							Phone:          invoicepb.Client.Phone,
							PostalCode:     invoicepb.Client.PostalCode,
							Address:        invoicepb.Client.Address,
						},
						CreatedAt: invoicepb.CreatedAt.AsTime(),
						UpdatedAt: invoicepb.UpdatedAt.AsTime(),
					}
					invoices = append(invoices, invoice)
				}
				mock_invoiceUsecase.EXPECT().GetInvoicesByDateRange(tt.request.StartDate.AsTime().String(), tt.request.EndDate.AsTime().String()).
					Return(invoices, nil)

			}

			s := &InvoiceServer{
				invoiceUsecase: mock_invoiceUsecase,
				companyUseCase: mock_companyUseCase,
				userUseCase:    mock_userUseCase,
			}
			got, err := s.GetInvoicesByDateRange(context.Background(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvoiceServer.GetInvoicesByDateRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want, protocmp.Transform()); diff != "" {
				t.Errorf("InvoiceServer.CreateInvoice() diff = %v", diff)
			}
		})
	}
}
