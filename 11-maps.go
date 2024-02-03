package main

import "fmt"

func main() {
	people := map[string]int{
		"amir":    20,
		"fatemeh": 21,
		"hossein": 20,
		"farshad": 34,
	}

	fmt.Println(people)
	fmt.Println(people["amir"])

	// Loop inside map
	for key, value := range people {
		fmt.Printf("Key: %v. Value: %v.\n", key, value)
	}

	people["sahar"] = 22
	people["hossein"] = 19

	fmt.Println(people)
}
