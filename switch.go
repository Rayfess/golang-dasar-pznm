package main

import "fmt"

func main() {
name := "Anon"

switch name {
case "Ave":
	fmt.Println("Hello Ave")
case "Rod":
	fmt.Println("Hello Rod")
default:
	fmt.Println("Who Are You?")
}

// short statement
name = "Viviparokastamu"
switch length := len(name); length > 10{
case true :
	fmt.Println("Your name is too long")
case false :
	fmt.Println("Your name in correct length")
}

// without expression	
name = "En"
length := len(name)
switch  {
case length > 10:
	fmt.Println("Your name is too long")
case length < 3:
	fmt.Println("Your name is too short")
default:
	fmt.Println("Your name in correct length")
}
}