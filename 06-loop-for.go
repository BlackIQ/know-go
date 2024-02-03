package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Printf("X is %v\n", x)
		x++
	}

	for y := 0; y < 5; y++ {
		fmt.Printf("Y is %v\n", y)
	}

	names := []string{"Amir", "Sahar", "Hossein", "Fatemeh"}

	for n := 0; n < len(names); n++ {
		fmt.Printf("Name is %v\n", names[n])
	}

	for index, value := range names {
		fmt.Printf("Index %v is: %v\n", index, value)
	}

	for _, value := range names {
		fmt.Printf("Name is %v\n", value)
	}
}
