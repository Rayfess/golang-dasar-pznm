package main

import "fmt"

// using alias / type declaration to simplified
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter ) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string  {
	if name == "Shit" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ave", spamFilter)

	// typecase2
	filter := spamFilter
	sayHelloWithFilter("Shit", filter)
}