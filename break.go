package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 { // stopped the loop if reached the statement
			break
		}
		fmt.Println("Looping for", i, "times")
	}
}