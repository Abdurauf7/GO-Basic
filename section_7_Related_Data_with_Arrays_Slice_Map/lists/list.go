package lists

import "fmt"

func Lists() {
	prices := []float64{1.98, 2.55}

	fmt.Println(prices)

	prices = append(prices, 1, 22, 55.22, 66.88, 99, 44, 500)

	discountPrices := []float64{101.99, 80.99, 20.59}

	prices = append(prices, discountPrices...) // concating two array elements with (element...)

	fmt.Println(prices)
}
