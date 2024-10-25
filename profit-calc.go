package main

import (
	"errors"
	"fmt"
)

func main() {
	revenue, error := getInput("Tell me your revenue")
	if error != nil {
		fmt.Print(error)
	}
	expenses, error := getInput("Tell me your expenses")
	if error != nil {
		fmt.Print(error)
	}

	tax, error := getInput("Tell me your tax rate")
	if error != nil {
		fmt.Print(error)
	}

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

func getInput(inputExpected string) (float64, error) {
	var inputValueProvided float64
	fmt.Println(inputExpected)
	fmt.Scan(&inputValueProvided)
	if inputValueProvided <= 0 {
		return 0, errors.New("Invalid input provided")

	}
	return inputValueProvided, nil
}
