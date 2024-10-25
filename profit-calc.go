package main

import "fmt"

func main() {
	var revenue float64
	fmt.Print("Tell me your revenue:")
	fmt.Scan(&revenue)
	var expenses float64
	fmt.Print("Tell me your expenses:")
	fmt.Scan(&expenses)
	var tax float64
	fmt.Print("Tell me your tax rate:")
	fmt.Scan(&tax)

	earnBeforeTax := revenue - expenses
	profit := earnBeforeTax * (1 - tax/100)

	fmt.Println("Your earnings before tax:")
	fmt.Println(earnBeforeTax)

	fmt.Println("Your earnings after tax:")
	fmt.Println(profit)

	fmt.Println("Ratio")
	fmt.Println(earnBeforeTax / profit)
}
