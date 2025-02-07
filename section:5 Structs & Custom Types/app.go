package main

import (
	"expample/section5/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Describing user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutpuUserDetails()
	appUser.ClearUserName()
	appUser.OutpuUserDetails()

	// Describing Admin

	adminUser, err := user.NewAdmin("test@mail.ru", "test12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	adminUser.OutpuUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
