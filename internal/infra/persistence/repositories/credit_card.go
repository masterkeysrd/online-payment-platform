package repositories

import (
	"github.com/masterkeysrd/online-payment-platform/internal/domain/creditcard"
	"gorm.io/gorm"
)

type creditCardRepository struct {
	db *gorm.DB
}

func NewCreditCardRepository(db *gorm.DB) creditcard.Repository {
	return &creditCardRepository{
		db: db,
	}
}

func (r *creditCardRepository) GetByPaymentMethod(merchantID, paymentMethodID string) (*creditcard.CreditCard, error) {
	var creditCard creditcard.CreditCard
	err := r.db.Where("merchant_id = ? AND payment_method_id = ?", merchantID, paymentMethodID).First(&creditCard).Error

	return &creditCard, err
}

func (r *creditCardRepository) Create(creditCard *creditcard.CreditCard) error {
	return r.db.Create(creditCard).Error
}
