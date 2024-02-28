package creditcard

import (
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/infra/id"
)

const (
	idPrefix = "cc_"
	idLength = 16
)

type Service interface {
	Create(request CreateCreditCardRequest) (CreateCreditCardResponse, error)
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

func (s *service) Create(request CreateCreditCardRequest) (CreateCreditCardResponse, error) {
	creditCard := CreditCard{
		ID:              generateId(),
		MerchantID:      request.Merchant,
		PaymentMethodId: request.PaymentMethod,
		Brand:           "visa",
		Number:          request.Number,
		Last4:           request.Number[len(request.Number)-4:],
		ExpMonth:        request.ExpMonth,
		ExpYear:         request.ExpYear,
		CVC:             request.CVC,
		Country:         "US",
		Created:         time.Now(),
	}

	err := s.repository.Create(&creditCard)
	if err != nil {
		return CreateCreditCardResponse{}, err
	}

	return CreateCreditCardResponse{
		ID:       creditCard.ID,
		Brand:    creditCard.Brand,
		Last4:    creditCard.Last4,
		ExpMonth: creditCard.ExpMonth,
		ExpYear:  creditCard.ExpYear,
		Country:  creditCard.Country,
		Created:  creditCard.Created.Unix(),
	}, nil
}

func generateId() string {
	return id.Generate(idPrefix, idLength)
}
