package services

import (
	"context"

	"calculator.com/internal/application/dto"
	"calculator.com/internal/application/repositories"
	"calculator.com/internal/domain"
	mathutil "calculator.com/internal/pkg/math"
)

type voucherService struct {
	repo repositories.VoucherRepository
}

var _ VoucherService = (*voucherService)(nil)

func (v voucherService) GetVoucherMaximumSubset(ctx context.Context, data dto.VoucherMaximumSubsetInput) (dto.VoucherMaximumSubsetOutput, error) {
	voucher := domain.NewVoucher(data.Voucher)
	if err := voucher.Validate(); err != nil {
		return dto.VoucherMaximumSubsetOutput{}, err
	}

	output := v.repo.GetVoucherMaximumSubset(ctx, data)
	output.VoucherSpent = mathutil.RoundFloatToTwoDecimalPlaces(output.VoucherSpent)

	return output, nil
}

func NewVoucherService(repo repositories.VoucherRepository) *voucherService {
	return &voucherService{
		repo: repo,
	}
}
