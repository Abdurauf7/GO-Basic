package practice

import "fmt"

type Products struct {
	title string
	id    string
	price float64
}

func Practice() {
	hobbies := [3]string{"reading", "gaming", "cycling"}

	combinedHobbies := hobbies[1:]

	fmt.Println("combinedHobbies", combinedHobbies)

	fmt.Println("hobbies", hobbies)

	slicedArray := hobbies[:2]

	hobbies[1] = "no-gaming"

	hobbies[2] = "no-cycling"

	fmt.Println("Re-slice", hobbies)

	fmt.Println(slicedArray)

	dynamicArr := []string{"learng go", "learn go rest api"}

	dynamicArr[1] = "different"

	dynamicArr = append(dynamicArr, "new-goal")

	dynamicProduct := []Products{{"The book", "1", 1000}, {"The book 2", "2", 200}}

	newProduct := Products{"Tablet", "3", 499.99}

	dynamicProduct = append(dynamicProduct, newProduct)
}
