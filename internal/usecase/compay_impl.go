package usecase

import (
	"github.com/tenling100/shiharaikun/internal/domain"
	"github.com/tenling100/shiharaikun/internal/repository"
)

type companyUsecase struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyUsecaseImpl(
	companyRepository repository.CompanyRepository,
) CompanyUsecase {
	return &companyUsecase{
		companyRepository: companyRepository,
	}
}

// CreateCompany creates a new company.
func (u *companyUsecase) CreateCompany(company *domain.Company) error {
	err := u.companyRepository.CreateCompany(company)
	if err != nil {
		return err
	}
	return nil
}
