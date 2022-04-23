package data

import (
	"context"
	"fmt"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/application/repositories"
	"calculator.com/pkg/httputils"
)

type productRepository struct {
	httpClient httputils.HTTPClient
}

var _ repositories.ProductRepository = (*productRepository)(nil)

func (p productRepository) GetProducts(ctx context.Context, IDs []string) ([]dto.Product, error) {
	var products []dto.Product

	doneCh := make(chan dto.Product)
	defer close(doneCh)
	errCh := make(chan error)
	defer close(errCh)

	for _, productID := range IDs {
		go func(productID string) {
			product, err := p.getProduct(ctx, productID)
			if err != nil {
				errCh <- err
				return
			}

			doneCh <- product
		}(productID)
	}

	for range IDs {
		select {
		case product := <-doneCh:
			products = append(products, product)
		case err := <-errCh:
			// When product not found, we don't want to return error
			// because we want to continue to get other products
			if err.Error() != "Product with id %s not found" {
				return nil, err
			}
			fmt.Println(err)
		}
	}

	return products, nil
}

func (p productRepository) getProduct(ctx context.Context, ID string) (dto.Product, error) {
	res, err := p.httpClient.Get(ctx, fmt.Sprintf("/items/%s", ID))
	if err != nil {
		return dto.Product{}, err
	}

	var product dto.Product
	if err = p.httpClient.ParseJSON(ctx, res, &product); err != nil {
		return dto.Product{}, err
	}

	return product, nil
}

func NewProductApiRepository(httpClient httputils.HTTPClient) *productRepository {
	return &productRepository{
		httpClient: httpClient,
	}
}
