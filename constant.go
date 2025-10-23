package main

import "fmt"
func main() {
	// Declare using Const, cannot redeclare value
	const age = 17
	fmt.Println(age)
	// age = 20, ERROR output

	// Multiple declare variables
	const (
		country = "Ireland"
		countryCode = 547238
	)
	fmt.Println(country)
	fmt.Println(countryCode)
}