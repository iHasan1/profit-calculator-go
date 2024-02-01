package main

import "fmt"

func main() {
	revenue, expenses, taxRate := getUserInput()

	earningsBeforeTax, profit, ratio := calcValues(revenue, expenses, taxRate)

	fmt.Println("Your Earnings Before Tax (EBT) is:", earningsBeforeTax)
	fmt.Println("Your Profit is:", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}

func getUserInput() (rev float64, exp float64, tR float64){
	fmt.Print("Enter your revenue: ")
	fmt.Scan(&rev)
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&exp)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&tR)
	return
}

func calcValues (rev, exp, tR float64) (ebt, profit, ratio float64){
	ebt = rev - exp
	profit = ebt * (1 - tR/100)
	ratio = ebt / profit
	return
}