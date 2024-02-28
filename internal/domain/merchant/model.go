package merchant

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
	CreatedAt  int64
}
