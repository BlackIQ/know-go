package main

import "fmt"

func main() {
	// Basic
	var strs = [3]string{"Amir", "Fatemeh", "Hossein"}
	// var nums [2]int = [2]int{128, 256}
	wired := [4]float64{1.1, 2.2, 3.3, 4.4}

	strs[2] = "Sahar"

	fmt.Println(strs)
	fmt.Println("To get length we use len(). Length of wired:", len(wired))

	// Appand
	points := []int{256, 512, 1024, 2048}
	points = append(points, 4096)

	fmt.Println(points)

	fmt.Println(points[2:4])
	fmt.Println(points[:3])
	fmt.Println(points[3:])
}
