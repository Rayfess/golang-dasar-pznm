package main

import "fmt"

func endApp()  {
	fmt.Println("End app")
	msg := recover()
	fmt.Println("panic Happens", msg) // handle the panic and return the message
}

func runApp(error bool)  {
	defer endApp()
	fmt.Println("OK 200")
	if error {
		panic("ERR")
	}
}

func main() {
	runApp(true) // will return error if true and end the app immediately if doesn't use recover()
	fmt.Println("Hello World")
}