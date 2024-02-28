package main

import (
	"github.com/masterkeysrd/online-payment-platform/api"
	"github.com/masterkeysrd/online-payment-platform/api/controllers"
)

func main() {
	// Controllers
	merchantController := controllers.NewMerchantController()
	customerController := controllers.NewCustomerController()

	server := api.NewServer()
	server.RegisterController("/api/v1/merchants", merchantController)
	server.RegisterController("/api/v1/customers", customerController)
	server.Run()
}
