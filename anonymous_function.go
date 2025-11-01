package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	blacklist := func (name string) bool  {
		return name == "Shit"
	}
	registerUser("Ave", blacklist)

	// or

	registerUser("Shit", func(name string) bool { // enter the name
		return name == "Shit" // if the value name is same as this blocked name, will respond the your blocked
	})
}