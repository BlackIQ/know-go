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

	fmt.Println()

	option, _ := getInput("[A]dd lesson, [R]eview report, [S]ave report, [E]xit: ", reader)

	switch option {
	case "A":
		fmt.Println()

		lesson, _ := getInput("Enter lesson: ", reader)
		score, _ := getInput("Enter score: ", reader)

		new, _ := strconv.ParseFloat(score, 64)

		u.updateLessons(lesson, new)

		fmt.Println("Lesson added.")

		getOption(u)
	case "R":
		fmt.Println()

		outout := u.format()

		fmt.Println(outout)

		getOption(u)
	case "S":
		data := []byte(u.format())

		outPath := fmt.Sprintf("files/%v.txt", u.name)

		err := os.WriteFile(outPath, data, 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println("Saved!\n")

	case "E":
		os.Exit(0)
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
