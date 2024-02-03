package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Amir"
	age := 18

	str := fmt.Sprintf("Name: %v. Age: %v.", name, age)

	data := []byte(str)

	err := os.WriteFile("lesson-19/test.txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("File saved.")
}
