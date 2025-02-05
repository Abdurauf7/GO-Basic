package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"


// Reading file
func getBalanceFromFile() float64{
	// if we don't use some argument we use _

	data, _ := os.ReadFile(accountBalanceFile)
	balanceConvertToText := string(data)
	balanceConvertToFloat, _ := strconv.ParseFloat(balanceConvertToText,64)
	return balanceConvertToFloat
}

// writing data into txt
func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile,[]byte(balanceText),0644)
}


func main() {
	var accountBalance = getBalanceFromFile()
	fmt.Println("\tWelcome to Go bank!!!")
	for {
		fmt.Println("*==================================*")
		fmt.Println("\tWhat do you want to do?")
		fmt.Println("\t1. Check balance")
		fmt.Println("\t2. Deposit money")
		fmt.Println("\t3. Withdraw money")
		fmt.Println("\t4. Exit")
		fmt.Println("*=================================*")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ") // Fixed quotes
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! Your amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!!!")
			fmt.Println("Thanks for choosing our bank!!!")
			// in switch case we use return;
			return 
			// in if statement we use break; 
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
