package main

import "fmt"

func main() {
	// const (
	// 	a = 5
	// 	b = 20
	// 	c = 100
	// 	z = 12
	// )

	// var d = a - b - c
	// var e = -a + b*c
	// var f = c / b
	// var g = z % a

	// println(d)
	// println(e)
	// println(f)
	// println(g)

// augmented assignment

	var i = 5
	i += 20 // i = i + 20	
	i /= 5 // i = i + 5	
	fmt.Println(i)

	//  unary operator
	var j = 1
	j++ // means var j +1
	fmt.Println(j)
}