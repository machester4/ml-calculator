package repositories

import (
	"context"

	"calculator.com/internal/application/dto"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, IDs []string) ([]dto.Product, error)
}
