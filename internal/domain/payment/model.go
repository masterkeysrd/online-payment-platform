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
	StatusReason    string
	TransactionID   string
	Created         time.Time
}
