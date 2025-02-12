package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Make is used to create a slice and efficiently allocate memory for optimal performance
	userName := make([]string, 2, 5) // -> []string - type of array| 2 - is pre defined array like adding two slots| 5 is length of array elements

	userName[0] = "Bob"

	userName = append(userName, "Max")
	userName = append(userName, "John")
	userName = append(userName, "Doe")
	userName = append(userName, "Smith")

	courseRating := make(floatMap, 3) // -> map[string]float64 - type of map| 3 - is length of map elements

	courseRating["Go"] = 4.5
	courseRating["Python"] = 4.7
	courseRating["Java"] = 4.2
	courseRating.output()

	// Loop through the array -- range is used to loop through the array key and value
	for index, value := range userName {
		fmt.Println(index, value)
	}

	// Loop through the map
	for key, value := range courseRating {
		fmt.Println(key, value)
	}
}
