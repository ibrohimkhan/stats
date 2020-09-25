package stats

import (
	"reflect"
	"github.com/ibrohimkhan/bank/v2/pkg/types"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment {
		{
			Amount: 1,
			Category: "auto",
		},
		{
			Amount: 2,
			Category: "auto",
		},
		{
			Amount: 3,
			Category: "auto",
		},
		{
			Amount: 550000,
			Category: "fun",
		},
	}

	expected := map[types.Category]types.Money {
		"auto": 2,
		"fun": 550000,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}