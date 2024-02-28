package payment

import (
	"errors"
	"log"
	"time"

	"github.com/masterkeysrd/online-payment-platform/internal/domain/acquiringbank"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/paymentmethod"
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
	repository           Repository
	paymentMethodService paymentmethod.Service
	acquiringBankService acquiringbank.Service
}

type ServiceParams struct {
	Repository           Repository
	PaymentMethodService paymentmethod.Service
	AcquiringBankService acquiringbank.Service
}

func NewService(params ServiceParams) Service {
	return &service{
		repository:           params.Repository,
		paymentMethodService: params.PaymentMethodService,
		acquiringBankService: params.AcquiringBankService,
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
	log.Println("[payment-service] creating payment")

	payment := Payment{
		ID:              generateID(),
		MerchantID:      request.Merchant,
		CustomerID:      request.Customer,
		PaymentMethodID: request.PaymentMethod,
		Description:     request.Description,
		Amount:          request.Amount,
		Currency:        request.Currency,
		Status:          "pending",
		StatusReason:    "pending to send to acquiring bank",
		Created:         time.Now(),
	}

	err := s.repository.Create(&payment)

	if err != nil {
		log.Println("[payment-service] error creating payment", err)
		return CreatePaymentResponse{}, err
	}

	log.Println("[payment-service] getting payment method")
	paymentMethod, err := s.paymentMethodService.GetForPayment(paymentmethod.GetPaymentMethodForPaymentRequest{
		Merchant:      request.Merchant,
		PaymentMethod: request.PaymentMethod,
		Customer:      request.Customer,
	})

	if err != nil {
		log.Println("[payment-service] error getting payment method", err)
		return CreatePaymentResponse{}, errors.Join(errors.New("error getting payment method"), err)
	}

	log.Println("[payment-service] sending transaction to acquiring bank")
	response, err := s.acquiringBankService.CreateTransaction(acquiringbank.CreateTransactionRequest{
		Sender: paymentMethod.Account,
		// TODO: Replace with merchant account
		Recipient:   "acquiring_bank-mock",
		Description: payment.Description,
		Amount:      payment.Amount,
	})

	if err != nil {
		log.Println("[payment-service] error sending transaction to acquiring bank", err)

		payment.Status = "failed"
		payment.StatusReason = err.Error()

		if err := s.repository.Update(&payment); err != nil {
			log.Println("[payment-service] error updating payment", err)
			return CreatePaymentResponse{}, errors.Join(errors.New("error updating payment"), err)
		}
	}

	if response.ID != "" {
		payment.TransactionID = response.ID
		payment.StatusReason = "transaction sent to acquiring bank"

		if err := s.repository.Update(&payment); err != nil {
			log.Println("[payment-service] error updating payment", err)
			return CreatePaymentResponse{}, errors.Join(errors.New("error updating payment"), err)
		}
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
