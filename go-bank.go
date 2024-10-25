package main

import (
	"fmt"
	"os"
)

func main() {
	var balance float64 = 1000

	fmt.Println("Welcome to GO BANK!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance.")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	choice := getInput("Give me your choice")

	if choice == 1 {
		fmt.Println("Your balance is", balance)
	} else if choice == 2 {
		deposit := getInput("How much?")
		balance += deposit
		fmt.Println("Your new balance is", balance)

	} else if choice == 3 {
		withdraw := getInput("How much?")
		balance -= withdraw
		fmt.Println("Your new balance is", balance)

	} else if choice == 4 {
		fmt.Println("Thank you visiting GO BANK!")
		os.Exit(0)

	} else {
		fmt.Println("Unsupported choice. Supoorted ones are 1,2,3,4")
	}

}

func getInput(inputExpected string) float64 {
	var inputValueProvided float64
	fmt.Println(inputExpected)
	fmt.Scan(&inputValueProvided)
	return inputValueProvided
}
