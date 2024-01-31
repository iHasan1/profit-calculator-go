package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println("Your Earnings Before Tax (EBT) is:", earningsBeforeTax)
	fmt.Println("Your Profit is:", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}