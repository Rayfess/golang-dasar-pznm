package main

import "fmt"

type validationErr struct {
	Msg string
}

func (v *validationErr) Error() string {
	return v.Msg
}

type notFoundErr struct {
	Msg string
}

func (n *notFoundErr) Error() string {
	return n.Msg
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationErr{"Validation Error!"}
	}

	if id != "user" {
		return &notFoundErr{"Data not Found!"}
	}

	//OK200

	return nil
}

func main() {
	err := SaveData("user", nil)
	if err != nil {
		/*
		 if validationErr, ok := err.(*validationErr); ok {
		 	fmt.Println("validation err:", validationErr.Error())
		 } else if notFoundErr, ok := err.(*notFoundErr); ok{
		 fmt.Println("not found err:", notFoundErr.Error())
		 } else {
		 	fmt.Println("unknown error:", err.Error())
		 }
		*/

		switch finalErr := err.(type) {
		case *validationErr:
			 	fmt.Println("validation err:", finalErr.Error())
		case *notFoundErr:
			fmt.Println("not found err:", finalErr.Error())
		default:
			fmt.Println("unknown error:", err.Error())
		}

	} else {
		fmt.Println("OK 200")
	}
}