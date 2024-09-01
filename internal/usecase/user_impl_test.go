package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/tenling100/shiharaikun/internal/domain"
	mock_repository "github.com/tenling100/shiharaikun/internal/repository/mock"
)

func Test_userUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	tests := []struct {
		name    string
		input   *domain.User
		wantErr bool
	}{
		{
			name: "normal case",
			input: &domain.User{
				Email:    "test",
				Password: "test",
			},
			wantErr: false,
		},
		{
			name: "error",
			input: &domain.User{
				Email:    "test",
				Password: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
			u := NewUserUsecaseImpl(mockUserRepository)
			if tt.wantErr {
				// error
				mockUserRepository.EXPECT().CreateUser(tt.input).Return(fmt.Errorf("error"))
			} else {
				// no error
				mockUserRepository.EXPECT().CreateUser(tt.input).Return(nil)
			}
			if err := u.CreateUser(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
