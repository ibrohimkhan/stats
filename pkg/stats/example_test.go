package stats

import (
	"fmt"
	"github.com/ibrohimkhan/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment {
		{
			Amount: 9,
		},
		{
			Amount: 7,
		},
		{
			Amount: 5,
		},
	}

	fmt.Println(Avg(payments))

	// Outputs:
	// 7
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
		},
	}

	fmt.Println(TotalInCategory(payments, "auto"))
	fmt.Println(TotalInCategory(payments, "pc"))

	// Outputs:
	// 16
	// 5
}