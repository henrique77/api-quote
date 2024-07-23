package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/henrique77/api-quote/model"
	controllerModel "github.com/henrique77/api-quote/model/controller"
	"github.com/henrique77/api-quote/service"
)

type QuoteController interface {
	SaveQuotes(app *fiber.Ctx) error
	GetMetrics(app *fiber.Ctx) error
}

type quoteController struct {
	service service.QuoteService
}

func NewQuoteController(service service.QuoteService) QuoteController {
	return &quoteController{
		service: service,
	}
}

func (q *quoteController) SaveQuotes(c *fiber.Ctx) error {
	request := new(controllerModel.QuoteRequest)

	if err := c.BodyParser(&request); err != nil {
		c.Status(http.StatusBadRequest).JSON(model.NewError().BadRequest(err))
	}

	response, err := q.service.Save(request)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(model.NewError().InternalServer(err))
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (q *quoteController) GetMetrics(c *fiber.Ctx) error {
	params := new(controllerModel.MetricsRequest)

	if err := c.QueryParser(params); err != nil {
		c.Status(http.StatusBadRequest).JSON(model.NewError().BadRequest(err))
	}

	response, err := q.service.GetMetrics(params.LastQuotes)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(model.NewError().InternalServer(err))
	}

	return c.Status(http.StatusOK).JSON(response)
}
