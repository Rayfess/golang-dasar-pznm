package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	edo := Man{"Edo"}
	edo.Married()

	fmt.Println(edo)
}