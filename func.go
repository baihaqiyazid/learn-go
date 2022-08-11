package main

import "fmt"

func getName(name string) {
	fmt.Println("Hello", name)
}

func getPerson() (string, int)  {
	return "Yazid", 21
}

func main() {
	
	fullName := "Muhammad Yazid Baihaqi"
	getName(fullName)

	name, _ := getPerson()
	fmt.Println(name)

}