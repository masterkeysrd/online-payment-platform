package acquiringbank

type CreateTransactionRequest struct {
	Sender      string  `json:"sender"`
	Recipient   string  `json:"recipient"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type CreateTransactionResponse struct {
	ID          string  `json:"id"`
	Sender      string  `json:"sender"`
	Recipient   string  `json:"recipient"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type TransactionEvent struct {
	ID        string          `json:"id"`
	Event     string          `json:"event"`
	Data      TransactionData `json:"data"`
	Signature string          `json:"signature"`
}

type TransactionData struct {
	TransactionID     string `json:"transaction_id"`
	AuthorizationCode string `json:"authorization_code"`
	ResponseCode      string `json:"response_code"`
	ResponseMessage   string `json:"response_message"`
}
