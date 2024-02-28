package merchant

import "time"

type Merchant struct {
	ID         string
	Name       string
	Email      string
	Phone      string
	Address    string
	Country    string
	Currency   string
	Website    string
	WebhookUrl string
	ApiKey     string
	Created    time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
