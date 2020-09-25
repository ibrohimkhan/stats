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

func TestPeriodsDynamic_negative(t *testing.T) {
	period1 := map[types.Category]types.Money {
		"auto": 10,
		"food": 20,
	}

	period2 := map[types.Category]types.Money {
		"auto": 5,
		"food": 3,
	}

	expected := map[types.Category]types.Money {
		"auto": -5,
		"food": -17,
	}

	result := PeriodsDynamic(period1, period2)

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestPeriodsDynamic_positiove(t *testing.T) {
	period1 := map[types.Category]types.Money {
		"auto": 10,
		"food": 20,
	}

	period2 := map[types.Category]types.Money {
		"auto": 20,
		"food": 20,
	}

	expected := map[types.Category]types.Money {
		"auto": 10,
		"food": 0,
	}

	result := PeriodsDynamic(period1, period2)

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestPeriodsDynamic_missedInSecond(t *testing.T) {
	period1 := map[types.Category]types.Money {
		"auto": 10,
		"food": 20,
	}

	period2 := map[types.Category]types.Money {
		"food": 20,
	}

	expected := map[types.Category]types.Money {
		"auto": -10,
		"food": 0,
	}

	result := PeriodsDynamic(period1, period2)

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestPeriodsDynamic_missedInFirst(t *testing.T) {
	period1 := map[types.Category]types.Money {
		"auto": 10,
		"food": 20,
	}

	period2 := map[types.Category]types.Money {
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money {
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}

	result := PeriodsDynamic(period1, period2)

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}