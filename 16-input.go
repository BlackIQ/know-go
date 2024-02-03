package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')

	fmt.Printf("Youe name is: %v\n", name)
}
