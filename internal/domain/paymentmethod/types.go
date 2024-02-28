package paymentmethod

type CreatePaymentMethodRequest struct {
	Merchant string                  `json:"merchant_id"`
	Customer string                  `json:"customer"`
	Type     string                  `json:"type"`
	Card     CreateCreditCardRequest `json:"card"`
}

type CreatePaymentMethodResponse struct {
	ID       string                   `json:"id"`
	Customer string                   `json:"customer"`
	Type     string                   `json:"type"`
	Card     CreateCreditCardResponse `json:"card"`
	Created  int64                    `json:"created"`
}

type CreateCreditCardRequest struct {
	Number   string `json:"number"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	CVC      string `json:"cvc"`
}

type CreateCreditCardResponse struct {
	ID       string `json:"id"`
	Brand    string `json:"brand"`
	Last4    string `json:"last4"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Country  string `json:"country"`
}
