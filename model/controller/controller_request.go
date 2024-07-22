package model

import "github.com/go-playground/validator/v10"

type Address struct {
	Zipcode string `json:"zipcode" validate:"required"`
}

type Recipient struct {
	Address Address `json:"address"`
}

type Volume struct {
	Category      int     `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Price         float64 `json:"price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type QuoteRequest struct {
	Recipient *Recipient `json:"recipient"`
	Volumes   []*Volume  `json:"volumes"`
}

func (c *QuoteRequest) Validate() error {
	validator := validator.New()

	return validator.Struct(c)
}
