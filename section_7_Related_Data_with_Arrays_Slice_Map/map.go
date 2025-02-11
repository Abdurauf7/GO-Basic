package main

import "fmt"

func main() {
	websites := map[string]string{"Google": "https://google.com", "Amazon": "https://aws.com"} // [key]: value

	fmt.Println(websites)
	fmt.Println(websites["Google"]) // extracting element

	websites["LinkedIn"] = "https:// linkedin.com" // adding new element

	delete(websites, "Google") // deleting element

	fmt.Println(websites)
}
