package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter)  {
	nameFilter := filter(name)

	fmt.Println("Hello", nameFilter)
}

func spam(name string) string {
	if name == "Anjing"{
		return "..."
	}else {
		return name
	}
}

func main() {
	
	sayHello("Yazid", spam)
	sayHello("Anjing", spam)

}