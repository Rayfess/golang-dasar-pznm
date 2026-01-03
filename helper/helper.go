package helper

import "fmt"

var version = "1.0.0"      // cant access, because not capitalize first word
var Application = "Helper" // can be accessed by another path by import

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Access() any{ 
	fmt.Println(version)
	return sayGoodbye("Axe")
}