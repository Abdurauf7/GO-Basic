package main

import "fmt"

func mainApp() {

	result := add(1, 2)
	fmt.Println(result)
}

// Generic function

func add[T int | float64 | string](a, b T) T {
	return a + b
}
