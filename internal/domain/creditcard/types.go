package creditcard

type GetCreditCardRequest struct {
	Merchant      string `json:"merchant"`
	PaymentMethod string `json:"payment_method"`
}

type GetCreditCardResponse struct {
	ID       string `json:"id"`
	Brand    string `json:"brand"`
	Last4    string `json:"last4"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Country  string `json:"country"`
	Created  int64  `json:"created"`
}

type GetCreditCardForPaymentRequest struct {
	ID            string `json:"id"`
	Merchant      string `json:"merchant"`
	PaymentMethod string `json:"payment_method"`
}

type GetCreditCardForPaymentResponse struct {
	ID       string `json:"id"`
	Brand    string `json:"brand"`
	Last4    string `json:"last4"`
	Number   string `json:"number"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Country  string `json:"country"`
	Created  int64  `json:"created"`
}

type CreateCreditCardRequest struct {
	Merchant      string `json:"merchant"`
	PaymentMethod string `json:"payment_method"`
	Number        string `json:"number"`
	ExpMonth      string `json:"exp_month"`
	ExpYear       string `json:"exp_year"`
	CVC           string `json:"cvc"`
}

type CreateCreditCardResponse struct {
	ID       string `json:"id"`
	Brand    string `json:"brand"`
	Last4    string `json:"last4"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Country  string `json:"country"`
	Created  int64  `json:"created"`
}
