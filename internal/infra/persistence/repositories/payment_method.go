package repositories

import (
	"github.com/masterkeysrd/online-payment-platform/internal/domain/paymentmethod"
	"gorm.io/gorm"
)

type paymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) paymentmethod.Repository {
	return &paymentMethodRepository{
		db: db,
	}
}

func (r *paymentMethodRepository) Create(paymentMethod *paymentmethod.PaymentMethod) error {
	return r.db.Create(paymentMethod).Error
}
