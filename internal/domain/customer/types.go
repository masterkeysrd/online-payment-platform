package customer

type CreateCustomerRequest struct {
	MerchantKey string `json:"merchant_key"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Country     string `json:"country"`
}

type CreateCustomerResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Country string `json:"country"`
	Created int64  `json:"created"`
}
