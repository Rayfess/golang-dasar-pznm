package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Edo")
	fmt.Println(result)

	result1 := helper.Access()
	fmt.Println(result1)
	
	fmt.Println(helper.Application) // access modifier can use, because the first character var is capitalized
}