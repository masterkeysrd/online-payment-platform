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

func (r *merchantRepository) GetByApiKey(apiKey string) (merchant.Merchant, error) {
	var m merchant.Merchant
	err := r.db.Where("api_key = ?", apiKey).First(&m).Error

	if err != nil {
		return merchant.Merchant{}, err
	}

	return m, nil
}
