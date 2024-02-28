package merchant

type CreateMerchantRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Country    string `json:"country"`
	Currency   string `json:"currency"`
	Website    string `json:"website"`
	WebhookUrl string `json:"webhook_url"`
}

type CreateMerchantResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Country    string `json:"country"`
	Currency   string `json:"currency"`
	Website    string `json:"website"`
	WebhookUrl string `json:"webhook_url"`
	ApiKey     string `json:"api_key"`
	Created    int64  `json:"created"`
}
