package main

import "fmt"

const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

func main() {
	fromCurrency := inputSourceCurrency()
	value := inputAmount()
	toCurrency := inputTargetCurrency(fromCurrency)

	result := getConvertedValue(value, fromCurrency, toCurrency)

	fmt.Println(result)
}

func inputSourceCurrency() string {
	for {
		fmt.Println("Выберите начальную валюту (USD, EUR, RUB)")

		currency := readCurrency()
		if isValidCurrency(currency) {
			return currency
		}

		fmt.Println("Выбранной валюты не существует в программе, введите ещё раз!")
	}
}

func inputTargetCurrency(fromCurrency string) string {
	for {
		fmt.Println("Выберите конечную валюту " + getAvailableCurrenciesHint(fromCurrency))

		currency := readCurrency()

		if !isValidCurrency(currency) {
			fmt.Println("Выбранной валюты не существует в программе, введите ещё раз!")
			continue
		}

		if currency == fromCurrency {
			fmt.Println("Эта валюта уже выбрана как начальная, выберите другую")
			continue
		}

		return currency
	}
}

func inputAmount() float64 {
	for {
		fmt.Println("Введите сумму:")

		value, ok := readNumber()
		if ok {
			return value
		}

		fmt.Println("Введите корректное число!")
		clearInputBuffer()
	}
}

func readCurrency() string {
	var currency string
	fmt.Scan(&currency)
	return currency
}

func readNumber() (float64, bool) {
	var value float64
	_, err := fmt.Scan(&value)

	if err != nil {
		return 0, false
	}

	return value, true
}

func clearInputBuffer() {
	var trash string
	fmt.Scanln(&trash)
}

func isValidCurrency(currency string) bool {
	return currency == USD || currency == EUR || currency == RUB
}

func getConvertedValue(value float64, fromCurrency string, toCurrency string) float64 {
	const usdToEur float64 = 0.93
	const usdToRub float64 = 87.34

	currencyRates := map[string]map[string]float64{
		USD: {
			EUR: usdToEur,
			RUB: usdToRub,
		},
		EUR: {
			USD: 1 / usdToEur,
			RUB: usdToRub / usdToEur,
		},
		RUB: {
			USD: 1 / usdToRub,
			EUR: usdToEur / usdToRub,
		},
	}

	return value * currencyRates[fromCurrency][toCurrency]
}

func getAvailableCurrenciesHint(selectedCurrency string) string {
	switch selectedCurrency {
	case USD:
		return "(EUR, RUB)"
	case EUR:
		return "(USD, RUB)"
	case RUB:
		return "(USD, EUR)"
	}

	return ""
}
