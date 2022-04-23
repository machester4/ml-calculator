package handlers

import (
	"net/http"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/engine"
	"github.com/gin-gonic/gin"
)

func HandleVoucherMaxSubset(c *gin.Context, e engine.ServiceEngine) {
	vs := e.GetVoucherService()
	ps := e.GetProductService()

	type body struct {
		ProductsIDs []string `json:"item_ids"`
		Amount      float64  `json:"amount"`
	}
	var input body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get products prices
	products, err := ps.GetProductsPrices(c, input.ProductsIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get voucher subset
	vInput := dto.VoucherMaximumSubsetInput{
		Products: products,
		Voucher:  input.Amount,
	}
	output, err := vs.GetVoucherMaximumSubset(c, vInput)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item_ids": output.ProductsIDs,
		"amount":   output.VoucherSpent,
	})
}
