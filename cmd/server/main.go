package main

import (
	"github.com/masterkeysrd/online-payment-platform/api"
)

func main() {
	server := api.NewServer()
	server.Run()
}