package funcitons

import "fmt"

type typeTransformFn func(int) int

func mainFuncitons() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 1, 2}

	// passing function to funtion
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("Doubled", doubled)
	fmt.Println("Tripled", tripled)

	// returning funtion inside funtion
	transformerFn1 := getTransformerFunction(&moreNumbers)
	transformerFn2 := getTransformerFunction(&numbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println("transformedNumbers", transformedNumbers)
	fmt.Println("moreTransformedNumbers", moreTransformedNumbers)

}

func transformNumbers(numbers *[]int, transform typeTransformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) typeTransformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
