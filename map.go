package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Ave",
		"address": "Tokyo",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	// func map
	book:= make(map[string]string)
	book["title"] = "Hero Village"
	book["author"] = "Ave"
	book["deleteMe"] = "invalid"

	fmt.Println("before deleted",book)

	delete(book, "deleteMe")
	fmt.Println("after deleted",book)
}