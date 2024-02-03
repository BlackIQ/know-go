package main

import (
	"fmt"
)

type user struct {
	name    string
	lessons map[string]float64
}

func add(name string) user {
	u := user{
		name:    name,
		lessons: map[string]float64{},
	}

	return u
}

func (u user) updateLessons(name string, score float64) {
	u.lessons[name] = score
}

func (u user) format() string {
	output := ""
	var total float64 = 0

	commonLine := fmt.Sprintf("+ %-10v + %-6v +\n", "----------", "------")

	output += fmt.Sprintf("+ %-10v +\n", "-------------------")
	output += fmt.Sprintf("| %-19v |\n", u.name)
	output += fmt.Sprintf("+ %-10v +\n", "-------------------")

	output += fmt.Sprintf("\n")

	output += commonLine
	output += fmt.Sprintf("+ %-10v + %-6v +\n", "Lesson", "Score")
	output += commonLine

	for key, value := range u.lessons {
		output += fmt.Sprintf("| %-10v | %-6v |\n", key, value)

		total += value
	}

	output += commonLine

	var miangin float64 = total / float64(len(u.lessons))

	output += fmt.Sprintf("\n")

	output += fmt.Sprintf("+ %-10v +\n", "-------------------")
	output += fmt.Sprintf("| Total: %-12v |\n", miangin)
	output += fmt.Sprintf("+ %-10v +\n", "-------------------")

	return output
}
