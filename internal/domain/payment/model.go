package payment

import "time"

type Payment struct {
	ID              string
	MerchantID      string
	CustomerID      string
	PaymentMethodID string
	Description     string
	Amount          float64
	Currency        string
	Status          string
	Created         time.Time
}
