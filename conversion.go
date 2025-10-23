package main

import "fmt"

func main() {
	var i32 int32 = 32768
	var i64 int64 = int64(i32)
	var i16 int16 = int16(i32) // number overflow, recount to min val of int16

	fmt.Println(i64)
	fmt.Println(i32)
	fmt.Println(i16)


	var name = "Erika"
	var e = name[0] // return output 69, In the standard ASCII character set, the uppercase letter 'E' has a decimal value of 69.
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}