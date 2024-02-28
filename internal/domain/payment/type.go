package payment

type GetPaymentRequest struct {
	Merchant  string `json:"merchant"`
	PaymentID string `json:"payment_id"`
}

type GetPaymentResponse struct {
	ID            string  `json:"id"`
	Customer      string  `json:"customer"`
	Description   string  `json:"description"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	Created       int64   `json:"created"`
}

type CreatePaymentRequest struct {
	Merchant      string  `json:"merchant"`
	Customer      string  `json:"customer"`
	Description   string  `json:"description"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
}

type CreatePaymentResponse struct {
	ID            string  `json:"id"`
	Customer      string  `json:"customer"`
	Description   string  `json:"description"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status"`
	Created       int64   `json:"created"`
}
