package repositories

import (
	"github.com/masterkeysrd/online-payment-platform/internal/domain/customer"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) customer.Repository {
	return &customerRepository{db}
}

func (r *customerRepository) Create(customer *customer.Customer) error {
	return r.db.Create(customer).Error
}
