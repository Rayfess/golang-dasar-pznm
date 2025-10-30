package main

import "fmt"

func getFullname() (string, string) {
	return "Axe", "Vex"
}

func main() {
	// firstName, lastName := getFullname()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullname() // ignoring return value
	fmt.Println(firstName)
}