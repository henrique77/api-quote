package client

import (
	"encoding/json"
	"fmt"
	"testing"

	clientModel "github.com/henrique77/api-quote/model/client"
	controllerModel "github.com/henrique77/api-quote/model/controller"
	"github.com/stretchr/testify/require"
)

var controllerRequest = []byte(`{
	"recipient":{
	"address":{
	"zipcode":"01311000"
	}
	},
	"volumes":[
	{
	"category":7,
	"amount":1,
	"unitary_weight":5,
	"price":349,
	"sku":"abc-teste-123",
	"height":0.2,
	"width":0.2,
	"length":0.2
	},
	{
	"category":7,
	"amount":2,
	"unitary_weight":4,
	"price":556,
	"sku":"abc-teste-527",
	"height":0.4,
	"width":0.6,
	"length":0.15
	}
	]
}`)

func Test_quoteClient_GetQuotes(t *testing.T) {
	request := new(controllerModel.QuoteRequest)

	err := json.Unmarshal(controllerRequest, request)
	require.NoError(t, err)

	clientRequest := new(clientModel.ClientQuoteRequest)

	clientRequest.New(request)

	client := NewQuoteClient("", "")

	response, err := client.GetQuotes(clientRequest)
	require.NoError(t, err)

	fmt.Println(response)

	t.Fail()
}
