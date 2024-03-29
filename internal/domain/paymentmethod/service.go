package paymentmethod

import (
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/domain/creditcard"
	"github.com/masterkeysrd/online-payment-platform/internal/infra/id"
)

const (
	CardType = "card"
	idPrefix = "pm"
	idLength = 16
)

type Service interface {
	Get(request GetPaymentMethodRequest) (GetPaymentMethodResponse, error)
	GetForPayment(request GetPaymentMethodForPaymentRequest) (GetPaymentMethodForPaymentResponse, error)
	Create(request CreatePaymentMethodRequest) (CreatePaymentMethodResponse, error)
}

type service struct {
	repository        Repository
	creditCardService creditcard.Service
}

type ServiceParams struct {
	Repository        Repository
	CreditCardService creditcard.Service
}

func NewService(params ServiceParams) Service {
	return &service{
		repository:        params.Repository,
		creditCardService: params.CreditCardService,
	}
}

func (s *service) Get(request GetPaymentMethodRequest) (GetPaymentMethodResponse, error) {
	paymentMethod, err := s.repository.Get(request.Merchant, request.Customer, request.PaymentMethod)

	if err != nil {
		return GetPaymentMethodResponse{}, err
	}

	var card GetCreditCardResponse
	switch paymentMethod.Type {
	case CardType:
		card, err = s.getCreditCard(request.Merchant, paymentMethod.ID)

		if err != nil {
			return GetPaymentMethodResponse{}, err
		}
	}

	return GetPaymentMethodResponse{
		ID:       paymentMethod.ID,
		Customer: paymentMethod.CustomerID,
		Type:     paymentMethod.Type,
		Card:     card,
		Created:  paymentMethod.Created.Unix(),
	}, nil
}

func (s *service) GetForPayment(request GetPaymentMethodForPaymentRequest) (GetPaymentMethodForPaymentResponse, error) {
	paymentMethod, err := s.repository.Get(request.Merchant, request.Customer, request.PaymentMethod)

	if err != nil {
		return GetPaymentMethodForPaymentResponse{}, err
	}

	var account string
	switch paymentMethod.Type {
	case CardType:
		card, err := s.creditCardService.GetForPayment(creditcard.GetCreditCardForPaymentRequest{
			Merchant:      request.Merchant,
			PaymentMethod: paymentMethod.ID,
		})

		if err != nil {
			return GetPaymentMethodForPaymentResponse{}, err
		}

		account = card.Number
	}

	return GetPaymentMethodForPaymentResponse{
		ID:       paymentMethod.ID,
		Customer: paymentMethod.CustomerID,
		Type:     paymentMethod.Type,
		Account:  account,
	}, nil
}

func (s *service) Create(request CreatePaymentMethodRequest) (CreatePaymentMethodResponse, error) {
	paymentMethod := PaymentMethod{
		ID:         generateId(),
		MerchantID: request.Merchant,
		CustomerID: request.Customer,
		Type:       request.Type,
		Created:    time.Now(),
	}

	err := s.repository.Create(&paymentMethod)
	if err != nil {
		return CreatePaymentMethodResponse{}, err
	}

	var createdCard CreateCreditCardResponse
	switch request.Type {
	case CardType:
		createdCard, err = s.createCreditCard(request.Merchant, paymentMethod.ID, request.Card)

		if err != nil {
			return CreatePaymentMethodResponse{}, err
		}
	}

	return CreatePaymentMethodResponse{
		ID:       paymentMethod.ID,
		Customer: paymentMethod.CustomerID,
		Type:     paymentMethod.Type,
		Card:     createdCard,
		Created:  paymentMethod.Created.Unix(),
	}, nil
}

func (s *service) getCreditCard(merchant string, paymentMethod string) (GetCreditCardResponse, error) {
	card, err := s.creditCardService.Get(creditcard.GetCreditCardRequest{
		Merchant:      merchant,
		PaymentMethod: paymentMethod,
	})

	if err != nil {
		return GetCreditCardResponse{}, err
	}

	return GetCreditCardResponse{
		ID:       card.ID,
		Brand:    card.Brand,
		Last4:    card.Last4,
		ExpMonth: card.ExpMonth,
		ExpYear:  card.ExpYear,
		Country:  card.Country,
		Created:  card.Created,
	}, nil
}

func (s *service) createCreditCard(merchant string, paymentMethod string, request CreateCreditCardRequest) (CreateCreditCardResponse, error) {
	card, err := s.creditCardService.Create(creditcard.CreateCreditCardRequest{
		Merchant:      merchant,
		PaymentMethod: paymentMethod,
		Number:        request.Number,
		ExpMonth:      request.ExpMonth,
		ExpYear:       request.ExpYear,
		CVC:           request.CVC,
	})

	if err != nil {
		return CreateCreditCardResponse{}, err
	}

	return CreateCreditCardResponse{
		ID:       card.ID,
		Brand:    card.Brand,
		Last4:    card.Last4,
		ExpMonth: card.ExpMonth,
		ExpYear:  card.ExpYear,
		Country:  card.Country,
	}, nil
}

func generateId() string {
	return id.Generate(idPrefix, idLength)
}
