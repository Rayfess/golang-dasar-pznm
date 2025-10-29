package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // will print the odd number if divisible by 0
			continue
		}
		fmt.Println("Looping for", i, "times")
	}
}