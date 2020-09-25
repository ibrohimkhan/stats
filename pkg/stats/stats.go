package stats

import (
	"github.com/ibrohimkhan/bank/v2/pkg/types"
)

// Avg calculates average payment
func Avg(payments []types.Payment) types.Money {
	size := len(payments)
	
	sum := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += int(payment.Amount)
		}
	}

	return types.Money(sum / size)
}

// TotalInCategory total sum in category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += int(payment.Amount)
		}
	}

	return types.Money(sum)
}

// CategoriesAvg average payment by category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	countedItems := map[types.Category]int{}

	for _, payment := range payments {
		countedItems[payment.Category]++
		categories[payment.Category] = (categories[payment.Category] + payment.Amount) / types.Money(countedItems[payment.Category])
	}

	return categories
}