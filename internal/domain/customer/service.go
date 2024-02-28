package customer

import (
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
	"github.com/masterkeysrd/online-payment-platform/internal/infra/id"
)

const (
	idPrefix = "cus"
	idLength = 14
)

type Service interface {
	CreateCustomer(request CreateCustomerRequest) (CreateCustomerResponse, error)
}

type service struct {
	repository      Repository
	merchantService merchant.Service
}

type ServiceParams struct {
	Repository      Repository
	MerchantService merchant.Service
}

func NewService(params ServiceParams) Service {
	return &service{
		repository:      params.Repository,
		merchantService: params.MerchantService,
	}
}

func (s *service) CreateCustomer(request CreateCustomerRequest) (CreateCustomerResponse, error) {
	merchantID, err := s.merchantService.GetIDByApiKey(request.MerchantKey)

	if err != nil {
		return CreateCustomerResponse{}, err
	}

	customer := Customer{
		ID:         generateId(),
		MerchantID: merchantID,
		Name:       request.Name,
		Email:      request.Email,
		Phone:      request.Phone,
		Address:    request.Address,
		Country:    request.Country,
		Created:    time.Now(),
	}

	err = s.repository.Create(&customer)

	if err != nil {
		return CreateCustomerResponse{}, err
	}

	return CreateCustomerResponse{
		ID:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
		Country: customer.Country,
		Created: customer.Created.Unix(),
	}, nil
}

func generateId() string {
	return id.Generate(idPrefix, idLength)
}
