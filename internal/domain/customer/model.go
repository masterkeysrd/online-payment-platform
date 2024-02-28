package customer

import (
	"time"
)

type Customer struct {
	ID         string
	MerchantID string
	Name       string
	Email      string
	Phone      string
	Address    string
	Country    string
	Created    time.Time
}
