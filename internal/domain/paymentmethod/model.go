package paymentmethod

import "time"

type PaymentMethod struct {
	ID         string
	MerchantID string
	CustomerID string
	Type       string
	Created    time.Time
}
