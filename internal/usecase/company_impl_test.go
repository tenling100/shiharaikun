package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/tenling100/shiharaikun/internal/domain"
	mock_repository "github.com/tenling100/shiharaikun/internal/repository/mock"
)

func Test_companyUsecase_CreateCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name     string
		input    *domain.Company
		expected *domain.Company
		wantErr  bool
	}{
		{
			name: "normal case",
			input: &domain.Company{
				Name:           "test",
				Representative: "test",
				Phone:          "test",
				PostalCode:     "test",
				Address:        "test",
			},
			wantErr: false,
		},
		{
			name: "error",
			input: &domain.Company{
				Name:           "test",
				Representative: "test",
				Phone:          "test",
				PostalCode:     "test",
				Address:        "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCompanyRepository := mock_repository.NewMockCompanyRepository(ctrl)
			if tt.wantErr {
				// error
				mockCompanyRepository.EXPECT().CreateCompany(tt.input).Return(fmt.Errorf("error"))
			} else {
				// no error
				mockCompanyRepository.EXPECT().CreateCompany(tt.input).Return(nil)
			}
			u := NewCompanyUsecaseImpl(mockCompanyRepository)
			if err := u.CreateCompany(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("companyUsecase.CreateCompany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
