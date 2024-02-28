package repositories

import (
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) merchant.Repository {
	return &merchantRepository{db}
}

func (r *merchantRepository) Create(merchant *merchant.Merchant) error {
	return r.db.Create(merchant).Error
}
