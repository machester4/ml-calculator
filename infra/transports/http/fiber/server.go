package fiberserver

import (
	"calculator.com/configs"
	"calculator.com/infra/transports"
	"calculator.com/internal/engine"
	"github.com/gofiber/fiber/v2"

	fiberhanders "calculator.com/infra/transports/http/fiber/handlers"
)

type fiberServer struct {
	app    *fiber.App
	config configs.ServerConfig
}

var _ transports.Transport = (*fiberServer)(nil)

func (s fiberServer) InitializeDefaultHandlers(e engine.ServiceEngine) {
	s.app.Get("/", s.wrap(fiberhanders.HandleHealthCheck, e))
}

func (s fiberServer) InitializeVoucherHandlers(e engine.ServiceEngine) {
	s.app.Post("/", s.wrap(fiberhanders.HandleVoucherMaxSubset, e))
}

type handler func(c *fiber.Ctx, e engine.ServiceEngine) error

func (s fiberServer) wrap(h handler, e engine.ServiceEngine) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(c, e)
	}
}

func (s fiberServer) Start() error {
	return s.app.Listen(s.config.Port)
}

func New(c configs.ServerConfig) *fiberServer {
	return &fiberServer{
		app: fiber.New(fiber.Config{
			Prefork: false,
		}),
		config: c,
	}
}
