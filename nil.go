package main

import "fmt"

func NewMap(name string) map[string]string  {
	if name == "" {
		return nil
	}else {
		return map[string]string{
			"name": name,
		}
	}
}

func main()  {
	person := NewMap("Yazid")

	if person == nil {
		fmt.Println("Data Nil")
	}else {
		fmt.Println(person)
	}
} 