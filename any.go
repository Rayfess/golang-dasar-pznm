package main

import "fmt"

func sample() any { // any == empty interface
	// return 1
	// return true
	return "A"
	
}

func main() {
	var empty any = sample()
	fmt.Println(empty)
}