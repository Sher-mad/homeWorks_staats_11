package stats_test

import (
	"fmt"

	"github.com/Sher-mad/homeWorks_bank_11/pkg/bank/types"
	"github.com/Sher-mad/homeWorks_stats_11/pkg.bank/stats"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   200,
			Category: "Караоке",
		},
		{
			ID:       2,
			Amount:   300,
			Category: "Караоке",
		},
		{
			ID:       3,
			Amount:   400,
			Category: "Караоке",
		},
	}
	fmt.Println(stats.Avg(payments))
	// Output:
	// 300
}

func ExampleTotalCategory() {
	payments := []types.Payment{
		{
			ID:       10,
			Amount:   120,
			Category: "Kafe",
		},
		{
			ID:       11,
			Amount:   130,
			Category: "SuperMarket",
		},
		{
			ID:       12,
			Amount:   140,
			Category: "Kafe",
		},
		{
			ID:       14,
			Amount:   150,
			Category: "SuperMarket",
		},

		{
			ID:       14,
			Amount:   5,
			Category: "SuperMarket",
		},
	}
	sumPayments := stats.TotalCategory(payments, types.Category("SuperMarket"))
	fmt.Println(sumPayments)
	//Output:
	// 285
	


}
