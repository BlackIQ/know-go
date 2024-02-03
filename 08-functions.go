package main

import (
	"fmt"
	"math"
)

func sayHi(name string) {
	fmt.Printf("Hello, %v.\n", name)
}

func cycleOperation(names []string, operator func(string)) {
	for _, value := range names {
		operator(value)
	}
}

func circleArea(r float64) float64 {
	data := math.Pi * r * r

	return data
}

func main() {
	name := "Amirhossein"
	names := []string{"Amir", "Sahar", "Hossein", "Fatemeh"}

	sayHi(name)
	cycleOperation(names, sayHi)

	a := circleArea(10)

	fmt.Println(a)
}
