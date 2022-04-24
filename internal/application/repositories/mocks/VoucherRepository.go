// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "calculator.com/internal/application/dto"
	mock "github.com/stretchr/testify/mock"
)

// VoucherRepository is an autogenerated mock type for the VoucherRepository type
type VoucherRepository struct {
	mock.Mock
}

// GetVoucherMaximumSubset provides a mock function with given fields: ctx, data
func (_m *VoucherRepository) GetVoucherMaximumSubset(ctx context.Context, data dto.VoucherMaximumSubsetInput) dto.VoucherMaximumSubsetOutput {
	ret := _m.Called(ctx, data)

	var r0 dto.VoucherMaximumSubsetOutput
	if rf, ok := ret.Get(0).(func(context.Context, dto.VoucherMaximumSubsetInput) dto.VoucherMaximumSubsetOutput); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(dto.VoucherMaximumSubsetOutput)
	}

	return r0
}
