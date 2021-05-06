package stats_test

import (
	"fmt"
	"reflect"
	"testing"

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

//__________________________________________________________lection_12_CategoriesAvd______________________________________________________
func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment

	result := stats.CategoriesAvg(payments)
	if len(result) != 0 {
		t.Error("  Result len ==0")
	}
}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}

	result := stats.CategoriesAvg(payments)
	if len(result) != 0 {
		t.Error("  Result len ==0")
	}
}


func TestCategoriesAvg_Total(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "cafe", Amount: 1_000, Status: "OK"},
		{ID: 2, Category: "food", Amount: 2_000, Status: "OK"},
		{ID: 3, Category: "cafe", Amount: 3_000, Status: "OK"},
		{ID: 4, Category: "cafe", Amount: 4_000, Status: "OK"},
		{ID: 5, Category: "fun", Amount: 5_000, Status: "OK"},
	}

	expected := map[types.Category]types.Money{
		"cafe": 2_666,
		"food": 2_000,
		"fun":  5_000,
	}

	result := stats.CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}

}
