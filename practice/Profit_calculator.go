package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Enter revenue amount: ")

	expenses, err2 := getUserInput("Enter expenses: ")

	tax_rate, err3 := getUserInput("Enter tax_rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
		// panic(err1)
	}

	ebt, profit, ration := calculateFinancials(revenue, expenses, tax_rate)

	febt := formattedText("EBT", ebt)
	fp := formattedText("PROFIT", profit)
	fr := formattedText("RATIO", ration)

	fmt.Print(febt, fp, fr)

	writeCalculationToFile(ebt, profit, ration)
}

func writeCalculationToFile(ebt, profit, ration float64) {
	calculationResult := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRation: %.3f\n", ebt, profit, ration)
	os.WriteFile("calculation.txt", []byte(calculationResult), 0644)

}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("invalid value, please input again")

	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, tax_rate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func formattedText(text string, value float64) (formattedValue string) {
	formattedValue = fmt.Sprintf(text+": %.1f\n", value)
	return formattedValue
}
