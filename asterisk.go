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


	address2 = &Address{"Paris", "London", "here"} // changing and creating new struct only for address2 by using pointer (&)
	// *address2 = Address{"Paris", "London", "here"} // overwriting the data struct but still same referencing the pointer 


	fmt.Println(address1) 
	fmt.Println(address2) 

}