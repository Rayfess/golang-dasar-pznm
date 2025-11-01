package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	// declaring function as value
	sample := getGoodBye

	fmt.Println(sample("Vray"))
}