package main

import "fmt"

func main() {
	var he = "Steve"
	var she = "Alex"

	var result bool = he == she // return false
	var result2 bool = he != she // return true

	fmt.Println(result, result2)
}