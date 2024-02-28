package creditcard

import "time"

// TODO: The Number and CVC fields should be encrypted.
type CreditCard struct {
	ID              string
	MerchantID      string
	PaymentMethodId string
	Brand           string
	Number          string
	Last4           string
	ExpMonth        string
	ExpYear         string
	Country         string
	CVC             string
	Created         time.Time
}
