package main

import "fmt"

// import "math"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax - (revenue * taxRate/100)

	fmt.Print("Your Earnings Before Tax (EBT) is:")
	fmt.Print(earningsBeforeTax)
	fmt.Println()
	fmt.Print("Your Profit is: ")
	fmt.Print(profit)
}