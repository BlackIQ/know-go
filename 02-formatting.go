package main

import "fmt"

func main() {
	name := "Amir"
	age := 20
	pi := 3.14159265359

	// Print and put data
	fmt.Println("My name is", name, "and I have", age, "years old")
	fmt.Printf("My name is %v and I have %v years old\n", name, age)

	// Work with float
	fmt.Printf("Pi type is %T\n", pi)
	fmt.Printf("Pi is %f\n", pi)
	fmt.Printf("Pi just 2 after 0 is %0.2f\n", pi)

	// Save strings
	var saved string = fmt.Sprintf("Pi just 2 after 0 is %0.2f\n", pi)
	fmt.Printf("Saved is: %s", saved)
}
