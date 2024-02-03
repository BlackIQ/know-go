package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{434, 6676, 232, 8878, 323, 1212, 454, 66}
	strs := []string{"Amir", "Sahar", "Hossein", "Fatemeh"}

	sort.Ints(ints) // Sort numbers

	fmt.Println(ints)
	fmt.Println(len(ints))                    // Get length
	fmt.Println(sort.SearchInts(ints, 232))   // Tells you where it is
	fmt.Println(sort.SearchInts(ints, 23200)) // Not there, one more than index of last

	sort.Strings(strs)

	fmt.Println(strs)
	fmt.Println(len(strs))                          // Get length
	fmt.Println(sort.SearchStrings(strs, "Sahar"))  // Tells you where it is
	fmt.Println(sort.SearchStrings(strs, "Jarvis")) // Not there, one more than index of last

}
