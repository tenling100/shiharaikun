package usecase

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/tenling100/shiharaikun/api"
	"github.com/tenling100/shiharaikun/internal/config"
	"github.com/tenling100/shiharaikun/internal/domain"
	mock_repository "github.com/tenling100/shiharaikun/internal/repository/mock"
)

func Test_invoiceUsecaseImpl_CreateInvoice(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name     string
		invoice  *domain.InvoiceData
		company  *domain.Company
		client   *domain.Company
		expected *domain.InvoiceData
		wantErr  bool
	}{
		{
			name: "normal case",
			invoice: &domain.InvoiceData{
				CompanyID:     1,
				ClientID:      2,
				InvoiceAmount: 1000,
			},
			company: &domain.Company{
				ID: 1,
			},
			client: &domain.Company{
				ID: 2,
			},
			expected: &domain.InvoiceData{
				CompanyID:     1,
				ClientID:      2,
				InvoiceAmount: 1000,
				FeeRate:       0.1,
				TaxRate:       0.1,
				Fee:           100,
				Tax:           100,
				PaymentAmount: 1200,
				Company: &domain.Company{
					ID: 1,
				},
				Client: &domain.Company{
					ID: 2,
				},
			},
			wantErr: false,
		},
		{
			name: "Fee and Tax are 0.1",
			invoice: &domain.InvoiceData{
				CompanyID:     1,
				ClientID:      2,
				InvoiceAmount: 1,
			},
			company: &domain.Company{
				ID: 1,
			},
			client: &domain.Company{
				ID: 2,
			},
			expected: &domain.InvoiceData{
				CompanyID:     1,
				ClientID:      2,
				InvoiceAmount: 1,
				FeeRate:       0.1,
				TaxRate:       0.1,
				Fee:           0.1,
				Tax:           0.1,
				PaymentAmount: 2,
				Company: &domain.Company{
					ID: 1,
				},
				Client: &domain.Company{
					ID: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvoiceDataRepository := mock_repository.NewMockInvoiceDataRepository(ctrl)
			mockCompanyRepository := mock_repository.NewMockCompanyRepository(ctrl)
			mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
			mockEnv := &config.Env{
				FeeRate: 0.1,
				TaxRate: 0.1,
			}
			// GetCompanyByID
			mockCompanyRepository.EXPECT().GetCompanyByID(tt.invoice.CompanyID).Return(tt.company, nil)
			// GetClient
			mockCompanyRepository.EXPECT().GetCompanyByID(tt.invoice.ClientID).Return(tt.client, nil)
			// CreateInvoiceData
			mockInvoiceDataRepository.EXPECT().CreateInvoiceData(tt.expected).Return(nil)
			u := NewInvoiceUsecaseImpl(mockInvoiceDataRepository, mockCompanyRepository, mockUserRepository, mockEnv)
			if err := u.CreateInvoice(tt.invoice); (err != nil) != tt.wantErr {
				t.Errorf("invoiceUsecase.CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func Test_invoiceUsecaseImpl_GetInvoicesByDateRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name     string
		request  *pb.DateRangeRequest
		expected []domain.InvoiceData
		wantErr  bool
	}{
		{
			name: "normal case",
			request: &pb.DateRangeRequest{
				StartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				EndDate:   timestamppb.New(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expected: []domain.InvoiceData{
				{
					ID:            1,
					CompanyID:     1,
					ClientID:      2,
					InvoiceAmount: 1000,
					FeeRate:       0.1,
					TaxRate:       0.1,
					Fee:           100,
					Tax:           100,
					PaymentAmount: 1200,
					Company:       &domain.Company{ID: 1},
					Client:        &domain.Company{ID: 2},
					IssueDate:     time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:            2,
					CompanyID:     1,
					ClientID:      2,
					InvoiceAmount: 1000,
					FeeRate:       0.1,
					TaxRate:       0.1,
					Fee:           100,
					Tax:           100,
					PaymentAmount: 1200,
					Company:       &domain.Company{ID: 1},
					Client:        &domain.Company{ID: 2},
					IssueDate:     time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "no data",
			request: &pb.DateRangeRequest{
				StartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				EndDate:   timestamppb.New(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "error",
			request: &pb.DateRangeRequest{
				StartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				EndDate:   timestamppb.New(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			expected: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvoiceDataRepository := mock_repository.NewMockInvoiceDataRepository(ctrl)
			mockCompanyRepository := mock_repository.NewMockCompanyRepository(ctrl)
			mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
			mockEnv := &config.Env{
				FeeRate: 0.1,
				TaxRate: 0.1,
			}
			if tt.wantErr {
				// GetInvoicesByDateRange
				mockInvoiceDataRepository.EXPECT().GetInvoicesByDateRange(tt.request.StartDate.String(), tt.request.EndDate.String()).
					Return(nil, fmt.Errorf("error"))
			} else {
				// GetInvoicesByDateRange
				mockInvoiceDataRepository.EXPECT().GetInvoicesByDateRange(tt.request.StartDate.String(), tt.request.EndDate.String()).
					Return(tt.expected, nil)
				for _, invoice := range tt.expected {
					// GetCompanyByID
					mockCompanyRepository.EXPECT().GetCompanyByID(invoice.CompanyID).Return(invoice.Company, nil)
					// GetClient
					mockCompanyRepository.EXPECT().GetCompanyByID(invoice.ClientID).Return(invoice.Client, nil)
				}
			}
			u := NewInvoiceUsecaseImpl(mockInvoiceDataRepository, mockCompanyRepository, mockUserRepository, mockEnv)
			var got []domain.InvoiceData
			var err error
			if got, err = u.GetInvoicesByDateRange(tt.request.StartDate.String(), tt.request.EndDate.String()); (err != nil) != tt.wantErr {
				t.Errorf("invoiceUsecase.GetInvoicesByDateRange() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("invoiceUsecase.GetInvoicesByDateRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}
