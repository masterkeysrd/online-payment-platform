package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	webhookURL string
}

type CreateTransactionRequest struct {
	Sender      string `json:"sender"`
	Recipient   string `json:"recipient"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type CreateTransactionResponse struct {
	ID          string `json:"id"`
	Sender      string `json:"sender"`
	Recipient   string `json:"recipient"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
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

func main() {
	router := gin.Default()

	config := Config{
		// webhookURL: os.Getenv("WEBHOOK_URL"),
		webhookURL: "http://localhost:8080/webhook",
	}

	router.POST("/transactions", createTransaction(config))

	router.Run(":8090")
}

func createTransaction(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		var transaction CreateTransactionRequest
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		time.Sleep(300 * time.Millisecond)

		event := simulateTransaction(transaction)
		go sendWebhook(config.webhookURL, event)

		c.JSON(201, CreateTransactionResponse{
			ID:          "123",
			Sender:      transaction.Sender,
			Recipient:   transaction.Recipient,
			Description: transaction.Description,
			Amount:      transaction.Amount,
		})
	}
}

func simulateTransaction(createTransaction CreateTransactionRequest) TransactionEvent {
	if createTransaction.Sender == "4242424242424242" {
		return TransactionEvent{
			ID:    "123",
			Event: "transaction",
			Data: TransactionData{
				TransactionID:     "123",
				AuthorizationCode: "123",
				ResponseCode:      "00",
				ResponseMessage:   "Approved",
			},
			Signature: "signature",
		}
	}

	if createTransaction.Sender == "4000000000000002" {
		return TransactionEvent{
			ID:    "123",
			Event: "transaction",
			Data: TransactionData{
				TransactionID:     "123",
				AuthorizationCode: "123",
				ResponseCode:      "05",
				ResponseMessage:   "Insufficient funds",
			},
			Signature: "signature",
		}
	}

	return TransactionEvent{
		ID:    "123",
		Event: "transaction",
		Data: TransactionData{
			TransactionID:     "123",
			AuthorizationCode: "123",
			ResponseCode:      "51",
			ResponseMessage:   "Insufficient funds",
		},
		Signature: "signature",
	}
}

func sendWebhook(url string, event TransactionEvent) {
	body, _ := json.Marshal(event)

	// Send the event to the webhook URL
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("Error sending webhook", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Webhook sent", response.Status)
}
