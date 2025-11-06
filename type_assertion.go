package main

import "fmt"

func random() any {
	return false
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)
	// var resultInt string = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown", value)
	}
}