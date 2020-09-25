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
	countedItems := map[types.Category]types.Money{}
	totalAmount := map[types.Category]types.Money{}

	for _, payment := range payments {
		countedItems[payment.Category]++
		totalAmount[payment.Category] += payment.Amount
	}

	for _, payment := range payments {
		categories[payment.Category] = totalAmount[payment.Category] / countedItems[payment.Category]
	}

	return categories
}

// PeriodsDynamic compares payments by category in periods
func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}

	for firstKey, firstValue := range first {
		value, ok := second[firstKey]
		if ok {
			result[firstKey] = value - firstValue

		} else {
			result[firstKey] = -firstValue
		}
	}

	for secondKey, secondValue := range second {
		vaule, ok := first[secondKey]
		if ok {
			result[secondKey] = secondValue - vaule
		
		} else {
			result[secondKey] = secondValue
		}
	}

	return result
}