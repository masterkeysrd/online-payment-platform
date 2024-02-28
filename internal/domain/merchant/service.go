package merchant

import "time"

type Service interface {
	CreateMerchant(request CreateMerchantRequest) (CreateMerchantResponse, error)
}

type service struct {
	merchants []Merchant
}

func NewService() Service {
	return &service{
		merchants: []Merchant{},
	}
}

func (s *service) CreateMerchant(request CreateMerchantRequest) (CreateMerchantResponse, error) {
	merchant := Merchant{
		ID:         "mrc_fa7d3f8d",
		Name:       request.Name,
		Email:      request.Email,
		Phone:      request.Phone,
		Address:    request.Address,
		Country:    request.Country,
		Currency:   request.Currency,
		Website:    request.Website,
		WebhookUrl: request.WebhookUrl,
		ApiKey:     "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
		CreatedAt:  time.Now().Unix(),
	}

	s.merchants = append(s.merchants, merchant)

	return CreateMerchantResponse(merchant), nil
}
