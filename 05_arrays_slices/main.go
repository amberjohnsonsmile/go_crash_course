package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	vegArr := [2]string{"Potato", "Carrot"}

	// Slices don't need to know number of items
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(vegArr)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
