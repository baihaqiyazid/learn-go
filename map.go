package main

import "fmt"

func main()  {
	
	person := map[string]string{
		"name": "Yazid",
		"city": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["city"])

	newPerson := make(map[string]string) //make map
	newPerson["name"] = "John"
	newPerson["city"] = "Bali"
	newPerson["age"] = "20"

	fmt.Println(newPerson)
	fmt.Println(newPerson["name"])
	fmt.Println(newPerson["city"])
	fmt.Println(newPerson["age"])
	fmt.Println(len(newPerson))

	delete(newPerson, "age") //DELETE
	fmt.Println(newPerson)

}