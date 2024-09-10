package main

import (
	"fmt"
	"math"
)

func main(){
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5;

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount) // & is pointer allows Scan to populate variable (not working with constant)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	
	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue :=  investmentAmount * math.Pow((1 + expectedReturnRate / 100),years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

// ** OUTPUT INFORMATION
	// fmt.Println("Future value: ",futureValue)
	fmt.Printf("Future value: %v\nFuture value (adjusted value): %v",futureValue,futureRealValue)
	// fmt.Println("Future value (adjusted value):",futureRealValue)
}