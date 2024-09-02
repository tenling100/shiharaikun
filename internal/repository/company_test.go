package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tenling100/shiharaikun/internal/domain"
)

func Test_companyRepository_CreateCompany(t *testing.T) {
	tests := []struct {
		name    string
		company *domain.Company
		wantErr bool
	}{
		{
			name: "normal case",
			company: &domain.Company{
				Name:           "test",
				Representative: "test",
				Phone:          "test",
				PostalCode:     "test",
				Address:        "test",
			},
			wantErr: false,
		},
		{
			name: "error case",
			company: &domain.Company{
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
			mockDB, mock, err := newMockDB()
			if err != nil {
				t.Errorf("failed to create mock db: %v", err)
			}
			defer mockDB.Callback()
			r := &companyRepository{
				db: mockDB,
			}
			mock.ExpectBegin()
			if tt.wantErr {
				mock.ExpectExec("INSERT INTO `companies`").
					WithArgs(
						tt.company.Name,
						tt.company.Representative,
						tt.company.Phone,
						tt.company.PostalCode,
						tt.company.Address,
						sqlmock.AnyArg(), // created_at
						sqlmock.AnyArg(), // updated_at
					).
					WillReturnError(fmt.Errorf("error"))
				mock.ExpectRollback()
			} else {
				mock.ExpectExec("INSERT INTO `companies`").
					WithArgs(
						tt.company.Name,
						tt.company.Representative,
						tt.company.Phone,
						tt.company.PostalCode,
						tt.company.Address,
						sqlmock.AnyArg(), // created_at
						sqlmock.AnyArg(), // updated_at
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}
			if err := r.CreateCompany(tt.company); (err != nil) != tt.wantErr {
				t.Errorf("companyRepository.CreateCompany() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Ensure all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %v", err)
			}

		})
	}
}

func Test_companyRepository_GetCompanyByID(t *testing.T) {
	tests := []struct {
		name      string
		requestID uint
		want      *domain.Company
		wantErr   bool
	}{
		{
			name: "normal case",
			want: &domain.Company{
				ID:             1,
				Name:           "test1",
				Representative: "test1",
				Phone:          "test1",
				PostalCode:     "test1",
				Address:        "test1",
			},
			wantErr: false,
		},
		{
			name:      "not found",
			requestID: 1,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := newMockDB()
			if err != nil {
				t.Errorf("failed to create mock db: %v", err)
			}
			defer mockDB.Callback()
			r := &companyRepository{
				db: mockDB,
			}
			// Set up expectations for the SQL query
			if tt.name == "not found" {
				mock.ExpectQuery("SELECT \\* FROM `companies` WHERE id = \\?").
					WithArgs(tt.requestID, sqlmock.AnyArg()).
					WillReturnError(fmt.Errorf("not found"))
			} else {
				mock.ExpectQuery("SELECT \\* FROM `companies` WHERE id = \\?").
					WithArgs(tt.requestID, sqlmock.AnyArg()).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "representative", "phone", "postal_code", "address"}).
							AddRow(tt.want.ID, tt.want.Name, tt.want.Representative, tt.want.Phone, tt.want.PostalCode, tt.want.Address),
					)
			}
			got, err := r.GetCompanyByID(tt.requestID)
			if (err != nil) != tt.wantErr {
				t.Errorf("companyRepository.GetCompanyByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("companyRepository.GetCompanyByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
