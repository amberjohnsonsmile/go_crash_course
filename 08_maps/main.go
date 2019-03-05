package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign keys/values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Heather"] = "heather@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add keys/values
	emails := map[string]string{"Bob": "bob@gmail.com", "Heather": "heather@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Heather"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
