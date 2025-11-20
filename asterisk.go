package main

import "fmt"

type Address struct {
	Country, Continent, City string
}

func main() {
	var address1 Address = Address{"Swiss", "chacha", "here"}
	var address2 *Address = &address1 // pointer

	address2.City = "there"

	// changed
	fmt.Println(address1)
	fmt.Println(address2)


	// address2 = &Address{"Paris", "London", "here"} // creating new struct cancel the changing of reference
	*address2 = Address{"Paris", "London", "here"} // creating new struct cancel the changing of reference
	fmt.Println(address1) // didnt changed
	fmt.Println(address2) // getting changed

}