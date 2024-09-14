package main

import "fmt"

func main()  {
	var accountBalance = 1000.0
	fmt.Println("\tWelcome to Go bank!!!")
	for {
		fmt.Println("*==================================*")
		fmt.Println("\tWhat do you want to do ?")
		fmt.Println("\t1. Check balance")
		fmt.Println("\t2. Deposite money")
		fmt.Println("\t3. Withdraw money")
		fmt.Println("\t4. Exit")
		fmt.Println("*=================================*")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		
		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposite: ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)
			if depositeAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater then zero")
				continue
			}
			accountBalance += depositeAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater then zero")
				continue
			}
			if(withdrawAmount > accountBalance){
				fmt.Println("Invalid amount. You cant withdraw more than you have")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! Your amount:",accountBalance)
		}else {
			fmt.Println("Goodbye!!!")
			//return
			break;
		}
	}
	fmt.Println("Thanks for choosing our bank!!!")
	
}