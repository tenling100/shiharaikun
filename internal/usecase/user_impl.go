package usecase

import (
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecaseImpl(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

// CreateUser creates a new user.
func (u *userUsecase) CreateUser(user *domain.User) error {
	err := u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
