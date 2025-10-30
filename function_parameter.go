package main

import "fmt"

func whoAreYou(firstName string, lastname string, age int) {
	fmt.Println("Hello", firstName, lastname, "Your", age)
}

func main() {
whoAreYou("Ave", "Vaslry", 90)
}