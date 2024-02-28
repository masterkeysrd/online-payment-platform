package merchant

import "time"

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

	err := s.repository.Create(&merchant)

	if err != nil {
		return CreateMerchantResponse{}, err
	}

	return CreateMerchantResponse(merchant), nil
}
