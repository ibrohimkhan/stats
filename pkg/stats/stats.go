package stats

import (
	"github.com/ibrohimkhan/bank/pkg/types"
)

// Avg calculates average payment
func Avg(payments []types.Payment) types.Money {
	size := len(payments)
	
	sum := 0
	for _, payment := range payments {
		sum += int(payment.Amount)
	}

	return types.Money(sum / size)
}

// TotalInCategory total sum in category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0

	for _, payment := range payments {
		if payment.Category == category {
			sum += int(payment.Amount)
		}
	}

	return types.Money(sum)
}