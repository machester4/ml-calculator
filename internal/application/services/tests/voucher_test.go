package tests

import (
	"context"
	"testing"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/application/services"
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
