package main

import (
	"fmt"
	"strconv"
)

func main() {
	var int1 int = 10
	var float1 float64 = 3.14
	var string1 string = "20"

	fmt.Printf("Type of int1 is %T.\n", int1)
	fmt.Printf("Type of float1 is %T.\n", float1)
	fmt.Printf("Type of string1 is %T.\n", string1)

	fmt.Println()

	strToFloat, _ := strconv.ParseFloat("20.5", 64)
	strToInt, _ := strconv.ParseInt("82", 64, 64)

	fmt.Printf("Type of strToFloat is %T.\n", strToFloat)
	fmt.Printf("Type of strToInt is %T.\n", strToInt)

	fmt.Println()

	intToStr := string(12)

	fmt.Printf("Type of intToStr is %T.\n", intToStr)
}
