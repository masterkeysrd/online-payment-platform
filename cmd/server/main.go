package main

import (
	"github.com/masterkeysrd/online-payment-platform/api"
	"github.com/masterkeysrd/online-payment-platform/api/controllers"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
)

func main() {
	// Services
	merchantService := merchant.NewService()

	// Controllers
	merchantController := controllers.NewMerchantController(
		controllers.MerchantControllerParams{
			Service: merchantService,
		},
	)

	customerController := controllers.NewCustomerController()

	server := api.NewServer()
	server.RegisterController("/api/v1/merchants", merchantController)
	server.RegisterController("/api/v1/customers", customerController)
	server.Run()
}
