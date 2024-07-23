package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrique77/api-quote/client"
	"github.com/henrique77/api-quote/config"
	"github.com/henrique77/api-quote/config/database"
	"github.com/henrique77/api-quote/controller"
	"github.com/henrique77/api-quote/repository"
	"github.com/henrique77/api-quote/service"
)

type server struct {
	router     *fiber.App
	env        *config.Env
	controller controller.QuoteController
}

func New() *server {
	return &server{}
}

func (s *server) Config() *server {
	s.env = config.ReadEnvs()
	s.router = fiber.New()

	repository := repository.NewQuoteRepository(database.InitDB(s.env))
	quoteClient := client.NewQuoteClient("", "")
	quoteService := service.NewQuoteService(quoteClient, repository)
	s.controller = controller.NewQuoteController(quoteService)

	s.configRouts()

	return s
}

func (s *server) Start() {
	s.env.Port = ":" + s.env.Port
	s.router.Listen(s.env.Port)
}

func (s *server) configRouts() {
	quoteV1 := s.router.Group("/v1")
	quoteV1.Post("/quote", s.controller.SaveQuotes)
	quoteV1.Get("/quote", s.controller.GetMetrics)
}
