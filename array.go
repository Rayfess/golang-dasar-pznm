package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ahmad"
	names[1] = "Bryan"
	names[2] = "Cassie"

	fmt.Println(names[2])

	//  Inline Array Inserting
	var values = [3]int{ 
		10,20, // if the value is empty, it will fill 0 or empty string
	}
	fmt.Println(values)
	fmt.Println(values[2])

	var lulues = [...]int{ //set the inline depends on the value of inserting
		10,50,100,80,
	}
	fmt.Println(lulues)
}
