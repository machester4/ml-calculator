package fiberserver

import (
	"calculator.com/configs"
	"calculator.com/infra/messaging"
	"calculator.com/infra/transports"
	"calculator.com/internal/engine"
	"github.com/gofiber/fiber/v2"

	fiberhanders "calculator.com/infra/transports/http/fiber/handlers"
	fibermiddleware "calculator.com/infra/transports/http/fiber/middleware"
)

type fiberServer struct {
	app    *fiber.App
	config configs.ServerConfig
}

var _ transports.Transport = (*fiberServer)(nil)

func (s fiberServer) InitializeDefaultHandlers(e engine.ServiceEngine) {
	s.app.Get("/", s.wrapHander(fiberhanders.HandleHealthCheck, e))
}

func (s fiberServer) wrapHander(h handler, e engine.ServiceEngine) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(c, e)
	}
}

type handler func(c *fiber.Ctx, e engine.ServiceEngine) error

func (s fiberServer) InitializeVoucherHandlers(e engine.ServiceEngine, m messaging.Messaging) {
	s.app.Post("/calculator/coupon",
		s.wrapMiddleware(fibermiddleware.MiddlewareVoucherMaxSubset, m),
		s.wrapHander(fiberhanders.HandleVoucherMaxSubset, e),
	)
}

func (s fiberServer) wrapMiddleware(mi middleware, m messaging.Messaging) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return mi(c, m)
	}
}

type middleware func(c *fiber.Ctx, m messaging.Messaging) error

func (s fiberServer) Start() error {
	return s.app.Listen(s.config.Port)
}

func New(c configs.ServerConfig) *fiberServer {
	return &fiberServer{
		app:    fiber.New(),
		config: c,
	}
}
