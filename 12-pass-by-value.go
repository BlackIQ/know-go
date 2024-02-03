package main

import "fmt"

func updateName(name string) string {
	name = "Fatemeh"

	return name
}

func updateMenu(menu map[string]float64) {
	menu["pizza"] = 7.99
}

func main() {
	// Non pointer values
	// Strings, Ints, Floats, Booleans, Arrays, Structs
	name := "Amirhossein"

	// Pointer wrapper values
	// Slices, Maps, Functions
	menu := map[string]float64{
		"tea":   1.99,
		"coffe": 3.99,
	}

	fmt.Println(name)
	fmt.Println(menu)

	name = updateName(name)
	updateMenu(menu)

	fmt.Println(name)
	fmt.Println(menu)

}
