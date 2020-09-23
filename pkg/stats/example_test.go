package stats

import (
	"fmt"
	"github.com/ibrohimkhan/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment {
		{
			Amount: 9,
			Status: types.StatusOk,
		},
		{
			Amount: 9,
			Status: types.StatusOk,
		},
		{
			Amount: 5,
			Status: types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	// Outputs:
	// 9
}

func ExampleTotalInCategory() {
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
			Status: types.StatusFail,
		},
	}

	fmt.Println(TotalInCategory(payments, "auto"))
	fmt.Println(TotalInCategory(payments, "pc"))

	// Outputs:
	// 16
	// 0
}