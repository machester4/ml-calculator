package fiberhandlers

import (
	"calculator.com/internal/engine"
	"github.com/gofiber/fiber/v2"
)

func HandleHealthCheck(c *fiber.Ctx, e engine.ServiceEngine) error {
	return c.SendString("ok")
}
