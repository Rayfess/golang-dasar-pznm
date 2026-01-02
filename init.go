package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // blank identifier
	"fmt"
)

func main() {
	fmt.Println(database.GetDB())
}