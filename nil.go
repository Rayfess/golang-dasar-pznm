package main

import "fmt"

// nil must be used in data type such as: interface, func, map, slice, pointer, channel

// func wrong(name string) string {
// 	if name == "" {
// 		return nil //cannot use nil as string value in return statement, return ERR
// 	} else {
// 		return name
// 	}
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil 
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
data := NewMap("Ada") //fill the map to return the name or else return nil

if data == nil {
	fmt.Println("Empty Data")
} else {
	fmt.Println(data["name"])
}

}