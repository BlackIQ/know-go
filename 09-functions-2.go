package main

import (
	"fmt"
	"strings"
)

func returnTwo(str string) (string, string) {
	var init []string

	for _, value := range strings.Split(str, " ") {
		init = append(init, strings.ToUpper(value[:1]))
	}

	if len(init) == 2 {
		return init[0], init[1]
	} else if len(init) == 1 {
		return init[0], "-"
	}

	return fmt.Sprintf("Len of input is %v", len(init)), "-"
}

func main() {
	a1, b1 := returnTwo("hello amir")
	fmt.Println(a1, b1)

	a2, b2 := returnTwo("amir")
	fmt.Println(a2, b2)

	a3, b3 := returnTwo("this is a long one")
	fmt.Println(a3, b3)
}
