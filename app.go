package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	earningsBeforeTax, profit, ratio := calcValues(revenue, expenses, taxRate)

	fmt.Println("Your Earnings Before Tax (EBT) is:", earningsBeforeTax)
	fmt.Println("Your Profit is:", profit)
	fmt.Printf("Ratio: %.2f", ratio)
}

func getUserInput(infoText string) (userInput float64){
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return
}

func calcValues (rev, exp, tR float64) (ebt, profit, ratio float64){
	ebt = rev - exp
	profit = ebt * (1 - tR/100)
	ratio = ebt / profit
	return
}