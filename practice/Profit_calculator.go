package main

import (
	"fmt"
)

func main()  {
	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Enter revenue amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax_rate: ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate / 100)
	ration := ebt / profit
	
	fmt.Println("EBT: ",ebt)
	fmt.Println("PROFIT: ",profit)
	fmt.Println("RATIO: ",ration)

}