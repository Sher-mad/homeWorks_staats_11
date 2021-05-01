package stats

import "github.com/Sher-mad/homeWorks_bank_11/v2/pkg/bank/types"

//Avg рассчмиывает среднюю сумма платёжаю
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status !=types.StatusFail{
			sum += payment.Amount
		}
	}

	return sum / types.Money(len(payments))
}

//TotalIntCategory находит сумму покупок по определенной кадегории
func TotallnCategory(payments []types.Payment, category types.Category) types.Money {
	total := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category && payment.Amount > 0 {
			total += payment.Amount
		}
	}
	return total

}
