package main

import "fmt"

func main() {
	numbers := []int{1, 20, 31, 11}

	sum := sumup(1, 20, 31, 11)
	secondSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(secondSum)
}

// startingNum -> 1
// number will be []int{ 20, 31, 11 }

func sumup(startingNum int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
