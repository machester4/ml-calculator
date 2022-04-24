package ginmiddleware

import (
	"context"

	"calculator.com/infra/messaging"
	"github.com/gin-gonic/gin"
)

func MiddlewareVoucherMaxSubset(c *gin.Context, m messaging.Messaging) {
	type body struct {
		ProductsIDs []string `json:"item_ids"`
		Amount      float64  `json:"amount"`
	}
	var input body
	if err := c.ShouldBindJSON(&input); err == nil {
		go func() {
			for _, id := range input.ProductsIDs {
				m.Publish(context.Background(), "voucher-metric", id)
			}
		}()
	}

	c.Next()
}
