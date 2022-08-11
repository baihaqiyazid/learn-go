package main

import "fmt"

type Customer struct {
	Name, City string
	Age		   int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi",name,"My Name is", customer.Name)
}

func main() {
	
	yazid := Customer {
		Name: "Yazid",
		City: "Jakarta",
		Age: 21,
	}

	// fmt.Println(yazid)
	yazid.sayHi("Mateo")
}