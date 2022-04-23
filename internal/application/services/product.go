package services

import (
	"context"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/application/repositories"
)

type productService struct {
	pr repositories.ProductRepository
}

var _ ProductService = (*productService)(nil)

func (p productService) GetProductsPrices(ctx context.Context, IDs []string) ([]dto.Product, error) {
	return p.pr.GetProducts(ctx, IDs)
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{
		pr: productRepository,
	}
}
