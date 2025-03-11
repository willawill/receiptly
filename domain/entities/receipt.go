package entities

import (
	"time"
	"github.com/go-playground/validator/v10"
)

type Item struct {
	shortDescription string `json:"shortDescription"`
	price            string `json:"price"`
}

type Receipt struct {
	ID           string `json:"id"`
	retailer     string `json:"retailer" binding:"required"`
	purchaseDate time.Time `json:"date" binding:"required"`
	purchaseTime time.Time `json:"time" binding:"required"`
	items        []Item `json:"items" binding:"required"`
	total        string `json:"total" binding:"required,gt=0"`
}

func ValidateReceipt(receipt Receipt) error {
	validate := validator.New()
	return validate.Struct(receipt)
}
