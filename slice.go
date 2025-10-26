package main

import "fmt"

func main() {
	names := [...]string{"Adam", "Henry", "Chris", "Beast", "Osas", "Xavier"} // an array

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:] // shortened from var slice3 []string = names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// Append Slice
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice1 := days[5:] // Saturday, Sunday, (pointer in 5) select afterwards
	fmt.Println("default array",days)
	fmt.Println("select the index after the pointer",daySlice1)

	daySlice1[0] = "saturNew"
	daySlice1[1] = "sundNew"

	fmt.Println("change the index",daySlice1)
	fmt.Println("applying changed index to default array",days)

	daySlice2 := append(daySlice1, "weekNew") //adding Data to Slice, it will try to adding first, but the array has reached max capacity and then he will make new Array(daysNew)
	// daysNew =:= [...]string{"saturNew", "sundNew", "weekNew" }
	daySlice2[0] = "saturOld" // will overwrite from daySlice1, first index
		fmt.Println("select the array from daySlice1, replaced saturNew",daySlice2)
		fmt.Println("nothing changed apart from first intermission",days)

	// make New Slice
	var newSlice []string = make([]string, 2, 5) // will not make an new array if the capacity not overlap with length number
	newSlice[0] = "Eve"
	newSlice[1] = "Eve"
	// newSlice[2] = "Eve" // return error, must use append to add new data

	fmt.Println(newSlice)
	fmt.Println("lenght of newSlice1 is",len(newSlice))
	fmt.Println("capacity of newSlice1 is",cap(newSlice))

	newSlice2 := append(newSlice, "sample") // correctly change the length
	fmt.Println(newSlice2)
		fmt.Println("lenght of newSlice2 is",len(newSlice2))
	fmt.Println("capacity of newSlice2 is",cap(newSlice2))

	newSlice2[0] = "Axe" // overwrite first index
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Diffrence between Making Array and Slice
	thisArray := [...]int{1, 2, 3, 4, 5} // or this one too [5]int{1,2,3,4,5}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)

	
}