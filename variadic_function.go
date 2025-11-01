package main

import "fmt"

// variable argument shortcut to calling slice/array but in function
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
func main() {
	fmt.Println(sumAll(10, 20, 30, 50))

	// slice param, converting from slice to varagrs
	numbers := []int{10,20,30,40}
	fmt.Println(sumAll(numbers...))
}