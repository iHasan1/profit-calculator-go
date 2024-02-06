package main

import "fmt"
import "errors"
import "os"

func storeResults(ebt, profit, ratio float64){
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calcValues(revenue, expenses, taxRate)

	fmt.Printf("Your Earnings Before Tax (EBT) is: %.1f\n", earningsBeforeTax)
	fmt.Printf("Your Profit is: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
	storeResults(earningsBeforeTax, profit, ratio)
}

func getUserInput(infoText string) (float64, error){
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number.")
	}
	return userInput, nil
}

func calcValues (rev, exp, tR float64) (ebt, profit, ratio float64){
	ebt = rev - exp
	profit = ebt * (1 - tR/100)
	ratio = ebt / profit
	return
}