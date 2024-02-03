package main

import (
	"fmt"
	"math"
)

func sayHi(name string) {
	fmt.Printf("Hello, %v.\n", name)
}

func circleArea(r float64) float64 {
	data := math.Pi * r * r

	return data
}

func whatIsX() {
	fmt.Println(theX)
}
