package main

import "fmt"

func main() {
	// looping for
	// count := 1;
	// for count <= 10{
	// 	fmt.Println("Looped by", count, "Times")
	// 	count++
	// }
	// fmt.Println("Done")

	// // for with statement
	// for count := 1; count <= 20; count++ {
	// 	fmt.Println("Looped by", count, "Times")
	// }
	// fmt.Println("Done")

	// for range used in data collection (array, slice, map)

	// accessing data manual
	names := []string{"Ave", "Exa", "Typhu"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}