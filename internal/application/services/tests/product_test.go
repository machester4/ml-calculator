package tests

import (
	"context"
	"fmt"
	"testing"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/application/repositories/mocks"
	"calculator.com/internal/application/services"
)

// Test the product service.

// When error in product repository, should return error.
func TestProductService_GetProducts_ErrorInProductRepository_ShouldReturnError(t *testing.T) {
	// Arrange
	productRepository := &mocks.ProductRepository{}
	productService := services.NewProductService(productRepository)

	// Act
	productRepository.On("GetProducts", context.Background(), []string{}).Return([]dto.Product{}, fmt.Errorf("error"))
	_, err := productService.GetProductsPrices(context.Background(), []string{})

	// Assert
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
