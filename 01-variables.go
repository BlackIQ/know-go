package main

import "fmt"

func main() {
	// String
	var name string = "Amir"
	var nicname = "GNU_Amir"
	var job string

	fmt.Println(name, nicname, job)

	name = "Developer"

	job = name

	fmt.Println(name, nicname, job)

	home := "Tehran"

	fmt.Println(home)

	fmt.Println("--------------------")

	// Integer
	var age1 int = 8
	var age2 = 16
	var age3 int

	fmt.Println(age1, age2, age3)

	age3 = 32

	fmt.Println(age1, age2, age3)

	age4 := 64

	fmt.Println(age4)

	fmt.Println("--------------------")

	// Bits and Memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	var secoreOne float32 = 1.5
	var secoreTwo float64 = 893823928392.83928323209

	fmt.Println(secoreOne, secoreTwo)
}
