package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// creating function which need to add numbers into
	doubleFnc := createTransformerFunction(2)
	triple := createTransformerFunction(3)

	// anonymus function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	// using this function inside another function
	doubled := transformNumbers(&numbers, doubleFnc)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("Transformed:", transformed)
	fmt.Println("Doubled:", doubled)
	fmt.Println("Tripled:", tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformerFunction(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
