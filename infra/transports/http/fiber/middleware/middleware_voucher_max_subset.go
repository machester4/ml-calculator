package fibermiddleware

import (
	"context"

	"calculator.com/infra/messaging"
	"github.com/gofiber/fiber/v2"
)

func MiddlewareVoucherMaxSubset(c *fiber.Ctx, m messaging.Messaging) error {
	type body struct {
		ProductsIDs []string `json:"item_ids"`
		Amount      float64  `json:"amount"`
	}
	var input body
	if err := c.BodyParser(&input); err == nil {
		go func() {
			for _, id := range input.ProductsIDs {
				m.Publish(context.Background(), "voucher-metric", id)
			}
		}()
	}

	return c.Next()
}
