package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Struct embedding

type Admin struct {
	email    string
	password string
	User     // inheriting User struct
}

func (u User) OutpuUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("Invalid Input")
	}
	return &Admin{
		email,
		password,
		User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName, lastName, birthdate string) (*User, error) {
	// validation of user

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("Invalid input try again")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
