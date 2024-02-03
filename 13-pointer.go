package main

import "fmt"

func updateIt(name *string) {
	*name = "Fatemeh"
}

func main() {
	name := "Amirhossein"

	fmt.Println(name)
	fmt.Println("Memory address of name is:", &name)

	memory := &name

	fmt.Println("Memory address:", memory)
	fmt.Println("Value is:", *memory)

	updateIt(memory)

	fmt.Println(name)

}
