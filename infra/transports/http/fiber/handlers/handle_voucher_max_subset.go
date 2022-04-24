package fiberhandlers

import (
	"calculator.com/internal/application/dto"
	"calculator.com/internal/engine"
	"github.com/gofiber/fiber/v2"
)

func HandleVoucherMaxSubset(c *fiber.Ctx, e engine.ServiceEngine) error {
	vs := e.GetVoucherService()
	ps := e.GetProductService()

	type body struct {
		ProductsIDs []string `json:"item_ids"`
		Amount      float64  `json:"amount"`
	}
	var input body
	if err := c.BodyParser(&input); err != nil {
		return fiber.ErrBadRequest
	}

	// Get products prices
	products, err := ps.GetProductsPrices(c.Context(), input.ProductsIDs)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Get voucher subset
	vInput := dto.VoucherMaximumSubsetInput{
		Products: products,
		Voucher:  input.Amount,
	}
	output, err := vs.GetVoucherMaximumSubset(c.Context(), vInput)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(map[string]interface{}{
		"item_ids": output.ProductsIDs,
		"amount":   output.VoucherSpent,
	})
}
