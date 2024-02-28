package creditcard

import "time"

type CreditCard struct {
	ID              string
	PaymentMethodId string
	Brand           string
	Number          string
	Last4           string
	ExpMonth        string
	ExpYear         string
	Country         string
	Created         time.Time
}
