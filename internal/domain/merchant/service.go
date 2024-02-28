package merchant

import (
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/infra/id"
)

const (
	idPrefix        = "mrc"
	idLength        = 14
	secretKeyPrefix = "sk"
	secretKeyLength = 24
)

type Service interface {
	GetIDByApiKey(apiKey string) (string, error)
	Create(request CreateMerchantRequest) (CreateMerchantResponse, error)
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

func (s *service) Create(request CreateMerchantRequest) (CreateMerchantResponse, error) {
	merchant := Merchant{
		ID:         generateId(),
		Name:       request.Name,
		Email:      request.Email,
		Phone:      request.Phone,
		Address:    request.Address,
		Country:    request.Country,
		Currency:   request.Currency,
		Website:    request.Website,
		WebhookUrl: request.WebhookUrl,
		ApiKey:     generateSecretKey(),
		Created:    time.Now(),
	}

	// TODO: Validate request, and check currency == "USD"
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

func (s *service) GetIDByApiKey(apiKey string) (string, error) {
	merchant, err := s.repository.GetByApiKey(apiKey)

	if err != nil {
		return "", err
	}

	return merchant.ID, nil
}

func generateId() string {
	return id.Generate(idPrefix, idLength)
}

func generateSecretKey() string {
	return id.Generate(secretKeyPrefix, secretKeyLength)
}
