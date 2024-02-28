package id

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Generate(prefix string, length int) string {
	id, _ := gonanoid.Generate(alphabet, length)
	return fmt.Sprintf("%s_%s", prefix, id)
}
