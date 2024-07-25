package service

import (
	"errors"

	"github.com/henrique77/api-quote/client"
	"github.com/henrique77/api-quote/model"
	clientModel "github.com/henrique77/api-quote/model/client"
	controllerModel "github.com/henrique77/api-quote/model/controller"
	"github.com/henrique77/api-quote/repository"
)

type QuoteService interface {
	Save(request *controllerModel.QuoteRequest) ([]*model.Quote, error)
	GetMetrics(lastQuotes int) (*model.Metrics, error)
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

	err := s.validateRequestInfo(request)
	if err != nil {
		return nil, err
	}
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

func (s *quoteService) GetMetrics(lastQuotes int) (*model.Metrics, error) {
	return s.repository.GetMetrics(lastQuotes)
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

func (s *quoteService) validateRequestInfo(request *controllerModel.QuoteRequest) error {
	if request.Recipient == nil {
		return errors.New("request is nil")
	}

	if request.Recipient.Address.Zipcode == "" {
		return errors.New("zipcode is empty")
	}

	if request.Volumes == nil {
		return errors.New("values is nil")
	}

	for _, v := range request.Volumes {
		if v.Category == 0 {
			return errors.New("category is empty")
		}
		if v.Amount == 0 {
			return errors.New("amunt is empty")
		}
		if v.UnitaryWeight == 0 {
			return errors.New("unitary wight is empty")
		}
		if v.Price == 0 {
			return errors.New("price is empty")
		}
		if v.Height == 0.0 {
			return errors.New("height is empty")
		}
		if v.Width == 0.0 {
			return errors.New("width is empty")
		}
		if v.Length == 0.0 {
			return errors.New("length is empty")
		}

	}

	return nil
}
