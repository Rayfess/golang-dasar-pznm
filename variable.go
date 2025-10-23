package main

import "fmt"
func main() {
	name := "Rhaef"
	fmt.Println(name)

	name = "Jacky"
	fmt.Println(name)

	// Multiple declare variables
	var (
		firstName = "Junior"
		lastName = "Downey"
	)
	fmt.Println(firstName, lastName)

	// Declare using Const, cannot redeclare value
	const age = 17
	fmt.Println(age)
}