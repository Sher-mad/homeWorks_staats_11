package stats

import "github.com/Sher-mad/homeWorks_bank_11/v2/pkg/bank/types"

//Avg рассчмиывает среднюю сумма платёжаю
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
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

//.._________________________________________________________________

//______________________________________Лекция_12________________________________________________________________________

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	sum := map[types.Category]types.Money{}
	result := map[types.Category]types.Money{}
	for _, payment := range payments {
		sum[payment.Category] += payment.Amount
		result[payment.Category]++
	}
	for i := 1; i < len(result); i++ {
		sum[payments[i].Category] /= result[payments[i].Category]
	}
	return sum

}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := make(map[types.Category]types.Money)
	// cycle := 0
	if len(first) < len(second) {
		// cycle = len(first)
		for key, value := range first {
			result[key] = value
		}
		for key, value := range second {
			result[key] -= value
			result[key] *= -1
		}
		for key, results := range result {
			if results > 0 {
				result[key] *= -1

			}
		}
	}
	if len(second) < len(first) {
		// cycle = len(second)
		for keySecond, valueSecond := range second {
			result[keySecond] = valueSecond
		}

		for key, value := range first {
			result[key] -= value
			result[key] *= -1
		}
		for key, results := range result {
			if results > 0 {
				result[key] *= -1

			}
		}

	}
	if len(first) == len(second) {
		// cycle = len(first)
		for key, value := range first {
			result[key] += value
		}
		for key, value := range second {
			result[key] -= value
			result[key] *= -1
		}
		for key, results := range result {
			if results > 0 {
				result[key] *= -1

			}
		}
	}

	return result

}
