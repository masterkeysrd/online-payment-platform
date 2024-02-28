package main

import (
	"github.com/masterkeysrd/online-payment-platform/api"
	"github.com/masterkeysrd/online-payment-platform/api/controllers"
)

func main() {
	// Controllers
	merchantController := controllers.NewMerchantController()

	server := api.NewServer()
	server.RegisterController("/api/v1/merchants", merchantController)
	server.Run()
}