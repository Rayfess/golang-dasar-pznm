package main

import (
	"errors"
	"fmt"
)

func Dividing(value int, divide int) (int, error) {
	if divide == 0 {
		return 0, errors.New("Divided by zERO")
	} else {
		return value / divide, nil
	}
}

func main() {
result, err := Dividing(1, 0)
if err == nil {
	fmt.Println("Result", result)
} else {
	fmt.Println("error:", err.Error())
}
}
