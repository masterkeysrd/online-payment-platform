package main

import (
	"github.com/masterkeysrd/online-payment-platform/api"
	"github.com/masterkeysrd/online-payment-platform/api/controllers"
	"github.com/masterkeysrd/online-payment-platform/api/middlewares"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/acquiringbank"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/creditcard"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/customer"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/payment"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/paymentmethod"
	"github.com/masterkeysrd/online-payment-platform/internal/infra/database"
	"github.com/masterkeysrd/online-payment-platform/internal/infra/persistence/repositories"
)

func main() {
	db := database.NewDatabase(&database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Name:     "online_payment_platform",
	})

	// Repositories
	merchantRepository := repositories.NewMerchantRepository(db)
	customerRepository := repositories.NewCustomerRepository(db)
	creditcardRepository := repositories.NewCreditCardRepository(db)
	paymentRepository := repositories.NewPaymentRepository(db)
	paymentMethodRepository := repositories.NewPaymentMethodRepository(db)

	// Services
	acquiringBankService := acquiringbank.NewService(acquiringbank.Config{
		URL: "http://localhost:8090",
	})

	merchantService := merchant.NewService(merchant.ServiceParams{
		Repository: merchantRepository,
	})

	customerService := customer.NewService(customer.ServiceParams{
		Repository:      customerRepository,
		MerchantService: merchantService,
	})

	creditCardService := creditcard.NewService(creditcard.ServiceParams{
		Repository: creditcardRepository,
	})

	paymentMethodService := paymentmethod.NewService(paymentmethod.ServiceParams{
		Repository:        paymentMethodRepository,
		CreditCardService: creditCardService,
	})

	paymentService := payment.NewService(payment.ServiceParams{
		Repository:           paymentRepository,
		PaymentMethodService: paymentMethodService,
		AcquiringBankService: acquiringBankService,
	})

	// Controllers
	merchantController := controllers.NewMerchantController(
		controllers.MerchantControllerParams{
			Service: merchantService,
		},
	)

	customerController := controllers.NewCustomerController(
		controllers.CustomerControllerParams{
			CustomerService:      customerService,
			PaymentMethodService: paymentMethodService,
		},
	)

	paymentController := controllers.NewPaymentController(
		controllers.PaymentControllerParams{
			PaymentService: paymentService,
		},
	)

	webhookController := controllers.NewWebhookController(
		controllers.WebhookControllerParams{},
	)

	server := api.NewServer()
	server.RegisterController("/api/v1/merchants", merchantController)
	server.RegisterController("/webhook", webhookController)
	server.RegisterControllerWithMiddleware("/api/v1/customers", customerController, middlewares.MerchantId(merchantService))
	server.RegisterControllerWithMiddleware("/api/v1/payments", paymentController, middlewares.MerchantId(merchantService))
	server.Run()
}
