package data

import (
	"context"

	"calculator.com/internal/application/dto"
	"github.com/machester4/greedy-approximation-algorithm/calculator"
)

type voucherRepository struct{}

var VoucherRepository = (*voucherRepository)(nil)

func (v voucherRepository) GetVoucherMaximumSubset(ctx context.Context, data dto.VoucherMaximumSubsetInput) dto.VoucherMaximumSubsetOutput {
	greedyItems := v.buildGreedyItems(data.Products)

	// Calculate maximum subset
	c := calculator.NewGreedyCalculator(greedyItems, data.Voucher)
	approximation := c.Calculate(ctx)

	productsIDs := v.getIDsFromGreedyItems(approximation.Items)

	return dto.VoucherMaximumSubsetOutput{
		ProductsIDs:  productsIDs,
		VoucherSpent: approximation.Amount,
	}
}

func (v voucherRepository) buildGreedyItems(products []dto.Product) []calculator.GreedyItem {
	greedyItems := make([]calculator.GreedyItem, len(products))
	for i, product := range products {
		greedyItems[i] = calculator.GreedyItem{
			ID:     product.ID,
			Amount: product.Price,
		}
	}

	return greedyItems
}

func (v voucherRepository) getIDsFromGreedyItems(greedyItems []calculator.GreedyItem) []string {
	ids := make([]string, len(greedyItems))
	for i, item := range greedyItems {
		ids[i] = item.ID
	}

	return ids
}

func NewVoucherRepository() *voucherRepository {
	return &voucherRepository{}
}
