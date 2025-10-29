package main

import "fmt"

func main() {
name := "Adavia"
if name == "Awe" {
	fmt.Println("hello Awe")
} else if name == "Ada"{
	fmt.Println("who  you?")	
} else {
	fmt.Println("who are you?")
}


// short statement
if length := len(name) ; length > 5 {
	fmt.Println("Names Too long")
} else {
	fmt.Println("Correctly")
}
}