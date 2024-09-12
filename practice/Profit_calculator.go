package main

import (
	"fmt"
)

func main()  {
	var revenue float64
	var expenses float64
	var tax_rate float64

	revenue = outputText("Enter revenue amount: ",revenue)
	expenses = outputText("Enter expenses: ",expenses)
	tax_rate = outputText("Enter tax_rate: ",tax_rate)

	ebt,profit,ration := calculateProfit(revenue,expenses,tax_rate)

	febt := formattedText("EBT",ebt)
	fp := formattedText("PROFIT",profit)
	fr := formattedText("RATIO",ration)

	fmt.Print(febt,fp,fr)
}

func outputText(text string, value float64) (float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculateProfit(revenue,expenses,tax_rate float64) (ebt float64, profit float64, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate / 100)
	ratio = ebt / profit

	return ebt,profit,ratio
}

func formattedText(text string, value float64) (formattedValue string){
	formattedValue = fmt.Sprintf(text + ": %.1f\n", value)
	return
}