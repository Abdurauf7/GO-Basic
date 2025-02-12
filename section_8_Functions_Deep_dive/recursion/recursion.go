package recursion

import "fmt"

// recursion
func recursion() {
	fact := factorial(5)

	fmt.Println(fact)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)

	// result := 1
	// for i := 1; i <= num; i++ {
	// 	result = result * i
	// }
	// return result
}
