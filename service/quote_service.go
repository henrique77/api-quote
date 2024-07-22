package service

import (
	"github.com/henrique77/api-quote/client"
	"github.com/henrique77/api-quote/model"
	clientModel "github.com/henrique77/api-quote/model/client"
	controllerModel "github.com/henrique77/api-quote/model/controller"
	"github.com/henrique77/api-quote/repository"
)

type QuoteService interface {
	Save(request *controllerModel.QuoteRequest) ([]*model.Quote, error)
}

type quoteService struct {
	quoteClient client.QuoteClient
	repository  repository.QuoteRepository
}

func NewQuoteService(quoteClient client.QuoteClient, repository repository.QuoteRepository) QuoteService {
	return &quoteService{
		quoteClient: quoteClient,
		repository:  repository,
	}
}

func (s *quoteService) Save(request *controllerModel.QuoteRequest) ([]*model.Quote, error) {
	clientRequest := new(clientModel.ClientQuoteRequest)

	clientRequest.New(request)

	cleintResponse, err := s.quoteClient.GetQuotes(clientRequest)
	if err != nil {
		return nil, err
	}

	quotes := s.readQuoteInfoFromClient(cleintResponse)

	if err := s.repository.Save(quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}

func (s *quoteService) readQuoteInfoFromClient(response *clientModel.ClientQuoteResponse) []*model.Quote {
	quotes := []*model.Quote{}

	for _, d := range response.Dispatchers {
		for _, o := range d.Offers {
			quotes = append(quotes, &model.Quote{
				Name:     o.Carrier.Name,
				Service:  o.Service,
				Deadline: s.extractDeadLine(&o.DeliveryTime),
				Price:    o.FinalPrice,
			})
		}
	}

	return quotes
}

func (s *quoteService) extractDeadLine(deliveryTime *clientModel.DeliveryTime) int {

	return deliveryTime.Days
}
