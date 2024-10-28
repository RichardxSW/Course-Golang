package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, tax float64

	revenue = getUserInput("Enter revenue: ")
	expenses = getUserInput("Enter expenses: ")
	tax = getUserInput("Enter tax: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, tax)

	printInfo(ebt, profit, ratio)
}

func calculateProfit(revenue, expenses, tax float64) (ebt, profit, ratio float64) {
	// var ebt, profit, ratio float64

	ebt = revenue - expenses
	profit = ebt * (1 - tax)
	ratio = ebt / profit

	return
}

func getUserInput(textInfo string) float64 {
	var userInput float64
	fmt.Print(textInfo)
	fmt.Scan(&userInput)
	return userInput
}

func printInfo(ebt, profit, ratio float64) {
	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}