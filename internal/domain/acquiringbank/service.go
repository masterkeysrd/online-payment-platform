package acquiringbank

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Service interface {
	CreateTransaction(request CreateTransactionRequest) (CreateTransactionResponse, error)
}

type service struct {
	config Config
}

func NewService(config Config) Service {
	return &service{
		config: config,
	}
}

func (s *service) CreateTransaction(request CreateTransactionRequest) (CreateTransactionResponse, error) {
	url := s.config.URL + "/transactions"

	var response CreateTransactionResponse
	err := postJSON(url, request, &response)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	return response, nil
}

func postJSON(url string, request interface{}, response interface{}) error {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	jsonRequest, err := json.Marshal(request)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonRequest))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(response)

	if err != nil {
		return err
	}

	return nil
}
