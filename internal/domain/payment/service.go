package payment

import (
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/infra/id"
)

const (
	idPrefix = "pay"
	idLength = 14
)

type Service interface {
	Get(GetPaymentRequest) (GetPaymentResponse, error)
	Create(CreatePaymentRequest) (CreatePaymentResponse, error)
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

func (s *service) Get(request GetPaymentRequest) (GetPaymentResponse, error) {
	payment, err := s.repository.Get(request.Merchant, request.PaymentID)

	if err != nil {
		return GetPaymentResponse{}, err
	}

	return GetPaymentResponse{
		ID:            payment.ID,
		Customer:      payment.CustomerID,
		Description:   payment.Description,
		Amount:        payment.Amount,
		Currency:      payment.Currency,
		PaymentMethod: payment.PaymentMethodID,
		Status:        payment.Status,
		Created:       payment.Created.Unix(),
	}, nil
}

func (s *service) Create(request CreatePaymentRequest) (CreatePaymentResponse, error) {
	payment := Payment{
		ID:              generateID(),
		MerchantID:      request.Merchant,
		CustomerID:      request.Customer,
		PaymentMethodID: request.PaymentMethod,
		Description:     request.Description,
		Amount:          request.Amount,
		Currency:        request.Currency,
		Status:          "pending",
		Created:         time.Now(),
	}

	err := s.repository.Create(&payment)

	if err != nil {
		return CreatePaymentResponse{}, err
	}

	return CreatePaymentResponse{
		ID:            payment.ID,
		PaymentMethod: payment.PaymentMethodID,
		Customer:      payment.CustomerID,
		Description:   payment.Description,
		Amount:        payment.Amount,
		Currency:      payment.Currency,
		Status:        payment.Status,
		Created:       payment.Created.Unix(),
	}, nil
}

func generateID() string {
	return id.Generate(idPrefix, idLength)
}
