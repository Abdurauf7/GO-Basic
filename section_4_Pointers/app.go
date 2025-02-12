package main

import "fmt"

func main() {
	age := 32 // regular variable

	// var agePointer *int

	// agePointer = &age // Here we get address of age variable from memory

	// second way to declare value of pointer
	agePointer := &age

	fmt.Println("AGE", *agePointer) // here we are get value from memory to show it

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult years", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // here overrding value of age 27
}
