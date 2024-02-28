package merchant

import (
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/infra/id"
)

type Service interface {
	CreateMerchant(request CreateMerchantRequest) (CreateMerchantResponse, error)
}

type service struct {
	repository Repository
}

type ServiceParams struct {
	Repository Repository
}

func NewService(params ServiceParams) Service {
	return &service{
		repository: params.Repository,
	}
}

func (s *service) CreateMerchant(request CreateMerchantRequest) (CreateMerchantResponse, error) {
	merchant := Merchant{
		ID:         id.Generate("mrc", 14),
		Name:       request.Name,
		Email:      request.Email,
		Phone:      request.Phone,
		Address:    request.Address,
		Country:    request.Country,
		Currency:   request.Currency,
		Website:    request.Website,
		WebhookUrl: request.WebhookUrl,
		ApiKey:     "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
		Created:    time.Now(),
	}

	err := s.repository.Create(&merchant)

	if err != nil {
		return CreateMerchantResponse{}, err
	}

	return CreateMerchantResponse{
		ID:         merchant.ID,
		Name:       merchant.Name,
		Email:      merchant.Email,
		Phone:      merchant.Phone,
		Address:    merchant.Address,
		Country:    merchant.Country,
		Currency:   merchant.Currency,
		Website:    merchant.Website,
		WebhookUrl: merchant.WebhookUrl,
		ApiKey:     merchant.ApiKey,
		Created:    merchant.Created.Unix(),
	}, nil
}
