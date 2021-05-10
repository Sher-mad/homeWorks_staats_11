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

//__________________________________________________________lection_12_PeriodsDynamic______________________________________________________

func TestPeriodsDynamic_nil(t *testing.T) {
	var first map[types.Category]types.Money
	var second map[types.Category]types.Money

	result := stats.PeriodsDynamic(first, second)

	if len(result) != 0 {
		fmt.Println("result len == 0")
	}

}
func TestPeriodsDynamic_empty(t *testing.T) {
	first := map[types.Category]types.Money{}
	second := map[types.Category]types.Money{}

	result := stats.PeriodsDynamic(first, second)

	if len(result) != 0 {
		fmt.Println("result len == 0", result)
	}

}
func TestPeriodsDynamic_first_result(t *testing.T) {
	first := map[types.Category]types.Money{

		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	expected := map[types.Category]types.Money{
		"auto": -5,
		"food": -17,
	}
	result := stats.PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}

}
func TestPeriodsDynamic_second_result(t *testing.T) {
	first := map[types.Category]types.Money{

		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto": 20,
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}
	result := stats.PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}

}

func TestPeriodsDynamic_third_result(t *testing.T) {
	first := map[types.Category]types.Money{

		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}
	result := stats.PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}

}

func TestPeriodsDynamic_fourth_result(t *testing.T) {
	first := map[types.Category]types.Money{

		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}
	expected := map[types.Category]types.Money{
		"auto":   0,
		"food":   -5,
		"mobile": -5,
	}
	result := stats.PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}

}
