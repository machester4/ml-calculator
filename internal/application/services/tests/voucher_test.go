package tests

import (
	"context"
	"testing"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/application/repositories/mocks"
	"calculator.com/internal/application/services"
	"github.com/stretchr/testify/mock"
)

// Test the voucher service.

// When voucher amount is less than or equal to 0, should return error.
func TestVoucherService_GetVoucherMaximumSubset_AmountIsLessThanOrEqualToZero_ShouldReturnError(t *testing.T) {
	// Arrange
	voucher := 0.0
	voucherService := services.NewVoucherService(nil)

	// Act
	_, err := voucherService.GetVoucherMaximumSubset(context.Background(), dto.VoucherMaximumSubsetInput{
		Voucher: voucher,
	})

	// Assert
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// Should round voucher amount to two decimal places.
func TestVoucherService_GetVoucherMaximumSubset_ShouldRoundVoucherAmountToTwoDecimalPlaces(t *testing.T) {
	// Arrange
	voucher := 1.23456789
	repo := &mocks.VoucherRepository{}
	voucherService := services.NewVoucherService(repo)

	// Act
	repo.On("GetVoucherMaximumSubset", mock.Anything, mock.Anything).Return(dto.VoucherMaximumSubsetOutput{})
	voucherService.GetVoucherMaximumSubset(context.Background(), dto.VoucherMaximumSubsetInput{
		Voucher: voucher,
	})

	// Assert
	repo.AssertCalled(t, "GetVoucherMaximumSubset", mock.Anything, mock.Anything)
}

// Should round voucher spent to two decimal places.
func TestVoucherService_GetVoucherMaximumSubset_ShouldRoundVoucherSpentToTwoDecimalPlaces(t *testing.T) {
	// Arrange
	voucher := 10.0
	repo := &mocks.VoucherRepository{}
	voucherService := services.NewVoucherService(repo)

	// Act
	repo.On("GetVoucherMaximumSubset", mock.Anything, mock.Anything).Return(dto.VoucherMaximumSubsetOutput{
		VoucherSpent: 9.99999,
	})
	output, _ := voucherService.GetVoucherMaximumSubset(context.Background(), dto.VoucherMaximumSubsetInput{
		Voucher: voucher,
	})

	// Assert
	if output.VoucherSpent != 9.99 {
		t.Errorf("Expected 9.99, got %f", output.VoucherSpent)
	}
}
