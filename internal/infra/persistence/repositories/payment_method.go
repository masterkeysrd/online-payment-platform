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

func (r *paymentMethodRepository) Get(merchantID, customerID, paymentMethodID string) (*paymentmethod.PaymentMethod, error) {
	var paymentMethod paymentmethod.PaymentMethod
	err := r.db.Where("merchant_id = ? AND customer_id = ? AND id = ?", merchantID, customerID, paymentMethodID).First(&paymentMethod).Error

	return &paymentMethod, err
}

func (r *paymentMethodRepository) Create(paymentMethod *paymentmethod.PaymentMethod) error {
	return r.db.Create(paymentMethod).Error
}
