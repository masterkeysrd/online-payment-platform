package main

import (
	"github.com/masterkeysrd/online-payment-platform/api"
	"github.com/masterkeysrd/online-payment-platform/api/controllers"
	"github.com/masterkeysrd/online-payment-platform/internal/domain/merchant"
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

	// Services
	merchantService := merchant.NewService(merchant.ServiceParams{
		Repository: merchantRepository,
	})

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
