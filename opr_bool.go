package main

import "fmt"

func main() {
	const examMin int = 80
	var studentValue int = 50

	var checkPassedStudent bool = studentValue >= examMin // return false, not qualify 
	
	fmt.Println(checkPassedStudent)
}