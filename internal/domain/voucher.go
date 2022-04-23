package domain

import "fmt"

type voucher struct {
	amount float64
}

func (v voucher) Amount() float64 {
	return v.amount
}

func (v voucher) Validate() error {
	if v.amount <= 0 {
		return fmt.Errorf("voucher amount must be greater than 0")
	}

	return nil
}

func NewVoucher(amount float64) *voucher {
	return &voucher{
		amount: amount,
	}
}
