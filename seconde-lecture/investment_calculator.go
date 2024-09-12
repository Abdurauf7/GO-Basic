package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main(){
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5;
	 
	investmentAmount = outputText("Enter investment amount: ",investmentAmount)
	expectedReturnRate = outputText("Enter Expected Return Rate: ",expectedReturnRate)
	years = outputText("Enter years: ",years)

	futureValue,futureRealValue := calculateFutureValue(investmentAmount,expectedReturnRate,years)

	// ** More formatted way display formatted text
	formattedFV := fmt.Sprintf("Future value: %.1f\n",futureValue) 
	formattedRFV := fmt.Sprintf("Future value (adjusted value): %.1f\n",futureRealValue)

// ** OUTPUT INFORMATION
	// fmt.Println("Future value: ",futureValue)
	// fmt.Printf("Future value: %.1f\nFuture value (adjusted value): %.1f",futureValue,futureRealValue)
	
// ** I can add more space in here
	// fmt.Printf(`Future value: %.1f Future value (adjusted value): %.1f`,futureValue,futureRealValue) 
	// fmt.Println("Future value (adjusted value):",futureRealValue)
	fmt.Print(formattedFV,formattedRFV)
}

// & is pointer allows Scan to populate variable (not working with constant)
func outputText(text string, value float64)(float64){
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

// by doing (fv float64,rfv float64) go create by him self variables in function
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64 ) (fv float64,rfv float64){
	fv = investmentAmount * math.Pow((1 + expectedReturnRate / 100),years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv,rfv
	// ** alternative way to return values
	// return
}