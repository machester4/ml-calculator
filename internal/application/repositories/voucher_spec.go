package repositories

import (
	"context"

	"calculator.com/internal/application/dto"
)

type VoucherRepository interface {
	GetVoucherMaximumSubset(ctx context.Context, data dto.VoucherMaximumSubsetInput) dto.VoucherMaximumSubsetOutput
}
