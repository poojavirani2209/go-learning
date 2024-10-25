package main

import "fmt"

func main() {
	revenue := getInput("Tell me your revenue")
	expenses := getInput("Tell me your expenses")
	tax := getInput("Tell me your tax rate")

	earnBeforeTax, profit := calculateProfit(revenue, expenses, tax)

	fmt.Println("Your earnings before tax:")
	fmt.Println(earnBeforeTax)

	fmt.Println("Your earnings after tax:")
	fmt.Println(profit)

	fmt.Println("Ratio")
	fmt.Println(earnBeforeTax / profit)
}

func calculateProfit(revenue float64, expenses float64, tax float64) (float64, float64) {
	earnBeforeTax := revenue - expenses
	profit := earnBeforeTax * (1 - tax/100)
	return earnBeforeTax, profit
}

func getInput(inputExpected string) float64 {
	var inputValueProvided float64
	fmt.Println(inputExpected)
	fmt.Scan(&inputValueProvided)
	return inputValueProvided
}
