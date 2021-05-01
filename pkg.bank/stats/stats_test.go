package stats_test

import (
	"fmt"

	"github.com/Sher-mad/homeWorks_bank_11/v2/pkg/bank/types"
	"github.com/Sher-mad/homeWorks_stats_11/v2/pkg.bank/stats"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   200,
			Category: "Караоке",
			Status:   "Ok",
		},
		{
			ID:       2,
			Amount:   300,
			Category: "Караоке",
			Status:   "Ok",
		},
		{
			ID:       3,
			Amount:   400,
			Category: "Караоке",
			Status:   "Ok",
		},
	}
	fmt.Println(stats.Avg(payments))
	// Output:
	// 300
}

func ExampleTotallnCategory() {
	payments := []types.Payment{
		{
			ID:       10,
			Amount:   120,
			Category: "Kafe",
			Status:   "Ok",
		},
		{
			ID:       11,
			Amount:   130,
			Category: "SuperMarket",
			Status:   "Ok",
		},
		{
			ID:       12,
			Amount:   140,
			Category: "Kafe",
			Status:   "Ok",
		},
		{
			ID:       14,
			Amount:   150,
			Category: "SuperMarket",
			Status:   "Ok",
		},

		{
			ID:       14,
			Amount:   5,
			Category: "SuperMarket",
			Status:   "Ok",
		},
	}
	sumPayments := stats.TotallnCategory(payments, types.Category("SuperMarket"))
	fmt.Println(sumPayments)
	//Output:
	// 285

}
