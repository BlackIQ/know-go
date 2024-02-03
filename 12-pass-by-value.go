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
	name := "Amirhossein"
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
