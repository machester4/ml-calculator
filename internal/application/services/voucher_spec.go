package services

import (
	"context"

	"calculator.com/internal/application/dto"
)

type VoucherService interface {
	GetVoucherMaximumSubset(ctx context.Context, data dto.VoucherMaximumSubsetInput) (dto.VoucherMaximumSubsetOutput, error)
}
