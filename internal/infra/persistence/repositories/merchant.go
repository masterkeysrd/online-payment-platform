package repositories

import (
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
)

type merchantRepository struct {
	merchants []merchant.Merchant
}

func NewMerchantRepository() merchant.Repository {
	return &merchantRepository{}
}

func (r *merchantRepository) Create(merchant *merchant.Merchant) error {
	r.merchants = append(r.merchants, *merchant)
	return nil
}
