package main

import (
	"fmt"
	"strings"
)

func main() {
	base := "Hello amir!"

	fmt.Println(strings.Contains(base, "amir"))
	fmt.Println(strings.ReplaceAll(base, "amir", "fatemeh"))
	fmt.Println(strings.ToUpper(base))
	fmt.Println(strings.ToLower(base))
	fmt.Println(strings.Index(base, "o"))
	fmt.Println(strings.Split(base, " "))
}
