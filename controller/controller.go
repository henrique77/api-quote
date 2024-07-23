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

// @Summary		Quote
// @Description	Route for receiving input data and generating a freight quote
// @Tags			quote
// @Accept			json
// @Produce		json
// @Param			request	body		model.QuoteRequest	true	"quote request"
// @Success		200		{array}		model.Quote
// @Failure		400		{object}	model.ControllerError
// @Failure		500		{object}	model.ControllerError
// @Router			/v1/quote [post]
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

// @Summary		Metrics
// @Description	Consult quote metrics
// @Tags			metrics
// @Accept			json
// @Produce		json
// @Param			last_quotes	query		int	false	"Number of quotes (descending order)"
// @Success		200			{array}		model.Metrics
// @Failure		400			{object}	model.ControllerError
// @Failure		500			{object}	model.ControllerError
// @Router			/v1/metrics [get]
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
