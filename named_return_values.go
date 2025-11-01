package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ave"
	middleName = "Crynoa" // if this deleted, will return default empty string
	lastName = "Vaxhu"

	return firstName, middleName, lastName
}
func main() {
	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a, b, c)
}