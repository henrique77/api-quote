package model

import (
	"fmt"
	"strconv"

	"github.com/henrique77/api-quote/config"
	controllerModel "github.com/henrique77/api-quote/model/controller"
)

type ClientQuoteRequest struct {
	Shipper        *Shipper      `json:"shipper"`
	Recipient      *Recipient    `json:"recipient"`
	Dispatchers    []*Dispatcher `json:"dispatchers"`
	Channel        string        `json:"channel,omitempty"`
	Filter         int           `json:"filter,omitempty"`
	Limit          int           `json:"limit,omitempty"`
	Identification string        `json:"identification,omitempty"`
	Reverse        bool          `json:"reverse,omitempty"`
	SimulationType []int         `json:"simulation_type"`
	Returns        *Returns      `json:"returns,omitempty"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number,omitempty"`
	StateInscription string `json:"state_inscription,omitempty"`
	Country          string `json:"country"`
	Zipcode          int    `json:"zipcode"`
}

type Volume struct {
	Amount        int     `json:"amount"`
	AmountVolumes int     `json:"amount_volumes,omitempty"`
	Category      string  `json:"category"`
	Sku           string  `json:"sku,omitempty"`
	Tag           string  `json:"tag,omitempty"`
	Description   string  `json:"description,omitempty"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Consolidate   bool    `json:"consolidate,omitempty"`
	Overlaid      bool    `json:"overlaid,omitempty"`
	Rotate        bool    `json:"rotate,omitempty"`
}

type Dispatcher struct {
	RegisteredNumber string    `json:"registered_number"`
	Zipcode          int       `json:"zipcode"`
	TotalPrice       float64   `json:"total_price,omitempty"`
	Volumes          []*Volume `json:"volumes"`
}

type Returns struct {
	Composition  bool `json:"composition,omitempty"`
	Volumes      bool `json:"volumes,omitempty"`
	AppliedRules bool `json:"applied_rules,omitempty"`
}

const (
	zipcodeDispatcher = 29161376
	country           = "BRA"
)

func (c *ClientQuoteRequest) New(quoteRequest *controllerModel.QuoteRequest) {
	zipCode, _ := strconv.Atoi(quoteRequest.Recipient.Address.Zipcode)

	registeredNumber := config.ReadEnvs().RegisteredNumber

	c.Shipper = &Shipper{
		RegisteredNumber: registeredNumber,
		Token:            config.ReadEnvs().TokenAPIFreteRapido,
		PlatformCode:     config.ReadEnvs().PlatformCode,
	}
	c.Recipient = &Recipient{
		Type:             1,
		RegisteredNumber: registeredNumber,
		Country:          country,
		Zipcode:          zipCode,
	}
	c.Dispatchers = []*Dispatcher{
		{
			RegisteredNumber: registeredNumber,
			Zipcode:          zipcodeDispatcher,
			Volumes:          c.fillVolumes(quoteRequest.Volumes),
		},
	}
	c.SimulationType = []int{0, 1}
}

func (c *ClientQuoteRequest) fillVolumes(volumes []*controllerModel.Volume) []*Volume {
	clientVolumes := []*Volume{}

	for _, v := range volumes {
		clientVolumes = append(clientVolumes, &Volume{
			Amount:        v.Amount,
			Category:      fmt.Sprint((v.Category)),
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
			UnitaryPrice:  v.Price,
			UnitaryWeight: v.UnitaryWeight,
		})
	}

	return clientVolumes
}
