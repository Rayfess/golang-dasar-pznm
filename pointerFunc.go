package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func setCountry(addr *Address, country string) {
	addr.Country = country
}

func main() {
	addr := &Address{} // var addr *Address = &Address{}
	setCountry(addr, "Japan")

	fmt.Println(addr)
}