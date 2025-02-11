
	=============== Array declaration ==================

	var productNames [4]string = [4]string{"laptop"}

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(prices[0])
	fmt.Println(productNames[0])

	productNames[2] = "car"

	fmt.Println(prices)
	fmt.Println(productNames)

	// Slice declaration

	featuredPrices := prices[1:3] // -> This will get the elements from index 1 to 3

	// featuredPrices := prices[:3] --> This will get the first 3 elements

	// featuredPrices := prices[1:] --> This will get the elements from index 1 to the end

	fmt.Println(featuredPrices)

	highlightedPrices := prices[:2]

	fmt.Println(highlightedPrices)

	fmt.Println(len(featuredPrices), cap(featuredPrices)) // -> len is length, cap is capacity
	highlightedPrices = highlightedPrices[:3]


======================= Adding new element ot array by append =================

prices := []float64{10.99, 8.99}

updatedPrices := append(prices, 5.99) // This will not work as append does not modify the original slice create one and add to it

prices = append(prices, 7.77) // adding new element to array

fmt.Println("prices", prices)

fmt.Println("updatedPrices", updatedPrices)
