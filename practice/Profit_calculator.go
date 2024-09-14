package main

import (
	"fmt"
)

func main()  {
	revenue := getUserInput("Enter revenue amount: ")
	expenses := getUserInput("Enter expenses: ")
	tax_rate := getUserInput("Enter tax_rate: ")

	ebt,profit,ration := calculateFinancials(revenue,expenses,tax_rate)

	febt := formattedText("EBT",ebt)
	fp := formattedText("PROFIT",profit)
	fr := formattedText("RATIO",ration)

	fmt.Print(febt,fp,fr)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue,expenses,tax_rate float64) (ebt float64, profit float64, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate / 100)
	ratio = ebt / profit

	return ebt,profit,ratio
}

func formattedText(text string, value float64) (formattedValue string){
	formattedValue = fmt.Sprintf(text + ": %.1f\n", value)
	return formattedValue
}