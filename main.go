package main

import "fmt"

func main() {
	const USD_TO_EUR float64 = 0.93
	const USD_TO_RUB float64 = 87.34
	const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
}

func getUserValue() float64 {
	var userValue float64

	fmt.Print("Введите значение: ")
	fmt.Scan(&userValue)

	return userValue
}

func getConvertedValue(value float64, fromCurrency string, toCurrency string) float64 {

}
