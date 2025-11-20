package main

import "fmt"

type Address struct {
	Country, Continent, City string
}

func main() {
	// by default go is using pass by value not by reference, so it will duplicate it to next target
	sample1 := Address{"USA", "Jakarta Selatan", "Amsterdam"}
	
	sample2 := sample1 // copy value
	sample2.City = "Greek"

	fmt.Println(sample1) // didnt change
	fmt.Println(sample2)


	// if using pointer you can pass data by reference, using operator (&)
	 var sample3 Address = Address{"USA", "Jakarta Selatan", "Amsterdam"}

	var sample4 *Address = &sample3 // pointer
	sample4.City = "Greek"

	//changed
	fmt.Println(sample3) 
	fmt.Println(sample4)
}