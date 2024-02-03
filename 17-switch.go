package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Age: ")
	data, _ := reader.ReadString('\n')

	age := strings.TrimSpace(data)

	switch age {
	case "20":
		fmt.Println("You are 20")
	case "25":
		fmt.Println("You are 25")
	case "35":
		fmt.Println("You are 35")
	default:
		fmt.Println("Default")
	}
}
