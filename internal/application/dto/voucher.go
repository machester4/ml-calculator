package dto

type VoucherMaximumSubsetInput struct {
	Products []Product
	Voucher  float64
}

type VoucherMaximumSubsetOutput struct {
	ProductsIDs  []string
	VoucherSpent float64
}
