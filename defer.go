package main

import "fmt"

func logging() {
	fmt.Println("Done Calling function")
}

func runApplication()  {
	defer logging() // will execute when the program ends even when error happends
	fmt.Println("Running the Application")
}
func main() {
runApplication()
}