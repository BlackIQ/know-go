package main

import "fmt"

func main() {
	age := 20

	if age < 20 {
		fmt.Println("You are less than 20")
	} else if age > 20 {
		fmt.Println("You are more than 20")
	} else {
		fmt.Println("You are 20")
	}
}
