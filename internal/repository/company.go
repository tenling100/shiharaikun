package repository

import (
	"gorm.io/gorm"

	"github.com/tenling100/shiharaikun/internal/domain"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *companyRepository {
	return &companyRepository{db: db}
}

// CreateCompany creates a new company.
func (r *companyRepository) CreateCompany(company *domain.Company) error {
	return r.db.Create(company).Error
}

// GetCompanyByID retrieves a company by its ID.
func (r *companyRepository) GetCompanyByID(id uint) (*domain.Company, error) {
	var company domain.Company
	err := r.db.Where("id = ?", id).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

// UpdateCompany updates a company.
func (r *companyRepository) UpdateCompany(company *domain.Company) error {
	return r.db.Save(company).Error
}
