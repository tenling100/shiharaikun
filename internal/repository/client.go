package repository

import (
	"gorm.io/gorm"

	"github.com/tenling100/shiharaikun/internal/domain"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *clientRepository {
	return &clientRepository{db: db}
}

// CreateClient creates a new client.
func (r *clientRepository) CreateClient(client *domain.Client) error {
	return r.db.Create(client).Error
}

// GetClientByID retrieves a client by its ID.
func (r *clientRepository) GetClientByID(id uint) (*domain.Client, error) {
	var client domain.Client
	err := r.db.Where("id = ?", id).First(&client).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&client).Association("BankAccounts").Find(&client.BankAccounts)
	if err != nil {
		return nil, err
	}
	err = r.db.Model(&client).Association("Invoices").Find(&client.Invoices)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// UpdateClient updates a client.
func (r *clientRepository) UpdateClient(client *domain.Client) error {
	return r.db.Save(client).Error
}
