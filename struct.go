package main

import "fmt"

//template data prototype data, knownly like object
type Customer struct { 
	Name, Address string
	Age           int
}

// struct method used in func as param
func (customer Customer) sayHello(name string) { 
	fmt.Println("Hello", name, "my name is", customer.Name, "and im", customer.Age)
}

func main() {
	// creating the data
	var ave Customer
	
	ave.Name = "Ave Brooklyn Ginger"
	ave.Address = "Brooklyn 67 Street"
	ave.Age = 35

	fmt.Println(ave)
	fmt.Println(ave.Age)

	// struct literal
	john := Customer{
		Name: "John Vocabulary",
		Address: "Sessame Street 99",
		Age: 29,
	}
	fmt.Println(john)
	fmt.Println(john.Address)


	knee := Customer {"Knee Ill", "Sewer West", 25}
	fmt.Println(knee)
	fmt.Println(knee.Name)


	ave.sayHello("Chris")
	knee.sayHello("Exa")
	john.sayHello("Valva")
}