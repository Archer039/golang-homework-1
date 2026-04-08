package main

import "fmt"

func main() {
	fromCurrency, toCurrency, value := getUserValue()
	result := getConvertedValue(value, fromCurrency, toCurrency)

	fmt.Println(result)
}

func getUserValue() (string, string, float64) {
	var fromCurrency, toCurrency string
	var value float64

	for {
		fmt.Println("Выберите начальную валюту (USD, EUR, RUB)")
		fmt.Scan(&fromCurrency)

		if fromCurrency == "USD" || fromCurrency == "EUR" || fromCurrency == "RUB" {
			break
		}

		fmt.Println("Выбранной валюты не существует в программе, введите ещё раз!")
	}

	for {
		fmt.Println("Введите сумму: ")
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println("Введите корректное число!")
			continue
		}
		break
	}

	for {
		fmt.Println("Выберите конечную валюту " + getAvailableCurrenciesHint(fromCurrency))
		fmt.Scan(&toCurrency)

		if fromCurrency == toCurrency {
			fmt.Println("Эта валюта уже выбранна как начальная, выберите другую")
			continue
		}

		if toCurrency == "USD" || toCurrency == "EUR" || toCurrency == "RUB" {
			break
		}
	}

	return fromCurrency, toCurrency, value
}

func getConvertedValue(value float64, fromCurrency string, toCurrency string) float64 {
	const UsdToEur float64 = 0.93
	const UsdToRub float64 = 87.34
	const EurToRub float64 = UsdToRub / UsdToEur

	var result float64

	switch {
	case fromCurrency == "USD" && toCurrency == "EUR":
		result = value * UsdToEur
	case fromCurrency == "EUR" && toCurrency == "USD":
		result = value / UsdToEur
	case fromCurrency == "EUR" && toCurrency == "RUB":
		result = value * EurToRub
	case fromCurrency == "RUB" && toCurrency == "EUR":
		result = value / EurToRub
	case fromCurrency == "USD" && toCurrency == "RUB":
		result = value * UsdToRub
	case fromCurrency == "RUB" && toCurrency == "USD":
		result = value / UsdToRub
	}

	return result
}

func getAvailableCurrenciesHint(selectedCurrency string) string {
	const USD string = "USD"
	const EUR string = "EUR"
	const RUB string = "RUB"

	var hint string

	switch selectedCurrency {
	case USD:
		hint = "(EUR, RUB)"
	case EUR:
		hint = "(USD, RUB)"
	case RUB:
		hint = "(USD, EUR)"
	}

	return hint
}
