package main

import "fmt"

type HasName interface{
	getName() string //Constructor return string
}

func sayHello(hasName HasName)  {
	fmt.Println("Hello", hasName.getName())
}

//Class
type Person struct {
	Name string
}

//Method
func (person Person) getName() string {
	return person.Name
}

func main() {
	
	yazid := Person{"Yazid"}

	sayHello(yazid)
}