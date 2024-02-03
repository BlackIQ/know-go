package main

import "fmt"

type person struct {
	name  string
	age   int
	score float64
}

func create(name string, age int, score float64) person {
	p := person{
		name:  name,
		age:   age,
		score: score,
	}

	return p
}

func (p person) format() string {
	str := fmt.Sprintf("| %-20v | %-8v | %-4v |", p.name, p.score, p.age)

	return str
}

func main() {
	amir := create("Amirhossein", 20, 15.08)
	fatemeh := create("Fatemeh", 20, 18)
	farshad := create("Farshad", 34, 20)

	fmt.Println(amir)

	fmt.Println("")

	commonLine := fmt.Sprintf("+ %-20v + %-8v + %-4v +\n", "--------------------", "--------", "----")

	fmt.Printf(commonLine)
	fmt.Printf("| %-20v | %-8v | %-4v |\n", "Name", "Score", "Age")
	fmt.Printf(commonLine)

	fmt.Println(amir.format())
	fmt.Println(fatemeh.format())
	fmt.Println(farshad.format())

	fmt.Printf(commonLine)
}
