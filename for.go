package main

import "fmt"

func main()  {
	
	var month = []string{
		"January",
		"February",
		"March",
		"April",
	}
	
	for i, val := range month {
		fmt.Println("index:", i, "value:", val)
	}

	newPerson := make(map[string]string) //make map
	newPerson["name"] = "John"
	newPerson["city"] = "Bali"
	newPerson["age"] = "20"

	for i, val := range newPerson {
		fmt.Println(i, val)
	}
}