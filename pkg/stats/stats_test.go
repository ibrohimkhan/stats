package stats

import (
	"reflect"
	"github.com/ibrohimkhan/bank/v2/pkg/types"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment {
		{
			Amount: 9,
			Category: "auto",
		},
		{
			Amount: 7,
			Category: "auto",
		},
		{
			Amount: 5,
			Category: "pc",
		},
	}

	expected := map[types.Category]types.Money {
		"auto": 8,
		"pc": 5,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}