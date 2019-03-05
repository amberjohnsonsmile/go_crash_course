package main

import "fmt"

func main() {
	a := 5
	// b is assigned to a pointer of a
	b := &a

	fmt.Println(a, b)
	// Use * to read val from address
	fmt.Println(*b)
	fmt.Printf("%T\n", b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
