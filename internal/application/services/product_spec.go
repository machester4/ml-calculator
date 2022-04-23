package services

import (
	"context"

	"calculator.com/internal/application/dto"
)

type ProductService interface {
	GetProductsPrices(ctx context.Context, IDs []string) ([]dto.Product, error)
}
