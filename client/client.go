package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/henrique77/api-quote/config"
	clientModel "github.com/henrique77/api-quote/model/client"
)

type QuoteClient interface {
	GetQuotes(request *clientModel.ClientQuoteRequest) (*clientModel.ClientQuoteResponse, error)
}

type quoteClient struct {
	host     string
	endpoint string
}

func NewQuoteClient(host string, endpoint string) QuoteClient {
	return &quoteClient{
		host:     host,
		endpoint: endpoint,
	}
}

func (c *quoteClient) GetQuotes(request *clientModel.ClientQuoteRequest) (*clientModel.ClientQuoteResponse, error) {
	urlApi := config.ReadEnvs().UrlAPIFreteRapido

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	requestHttp, err := http.NewRequest(http.MethodPost, urlApi, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	requestHttp.Header.Set("Content-Type", "application/json")

	responseHttp, err := http.DefaultClient.Do(requestHttp)
	if err != nil {
		return nil, err
	}
	defer responseHttp.Body.Close()

	if responseHttp.StatusCode != http.StatusOK {
		return nil, errors.New("Falha na requisição para o client")
	}

	responseBody, err := io.ReadAll(responseHttp.Body)

	if err != nil {
		return nil, err
	}

	var response *clientModel.ClientQuoteResponse

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}
