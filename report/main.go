package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error
}

func getOption(u user) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("[A]dd lesson, [S]ave report: ", reader)

	switch option {
	case "A":
		lesson, _ := getInput("Enter lesson: ", reader)
		score, _ := getInput("Enter score: ", reader)

		new, _ := strconv.ParseFloat(score, 64)

		u.updateLessons(lesson, new)

		getOption(u)
	case "S":
		fmt.Println("Saved!\n")

		outout := u.format()

		fmt.Println(outout)
	default:
		fmt.Println("Try again")
		getOption(u)
	}
}

func addUser() user {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter student name: ", reader)

	u := add(name)

	return u
}

func main() {
	u := addUser()

	getOption(u)
}
