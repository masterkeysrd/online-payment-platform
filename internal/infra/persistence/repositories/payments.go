package repositories

import (
	"github.com/masterkeysrd/online-payment-platform/internal/domain/payment"
	"gorm.io/gorm"
)

func NewPaymentRepository(db *gorm.DB) payment.Repository {
	return &paymentRepository{
		db: db,
	}
}

type paymentRepository struct {
	db *gorm.DB
}

func (r *paymentRepository) Get(merchantID string, paymentID string) (payment.Payment, error) {
	var p payment.Payment
	err := r.db.Where("merchant_id = ? AND id = ?", merchantID, paymentID).First(&p).Error
	return p, err
}

func (r *paymentRepository) Create(request *payment.Payment) error {
	return r.db.Create(request).Error
}
