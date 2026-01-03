package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := new(Address) //  var address1 *Address = &Address{}
	address2 := address1    //  var address2 *Address = address1

	address2.Country = "Japan"

	fmt.Println(address1)
	fmt.Println(address2)
}