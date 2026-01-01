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

	// study case to count how many are duplicate things in array
words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
wordCount := make(map[string]int)
for _, word := range words {
    wordCount[word]++
}
fmt.Println(wordCount)
}