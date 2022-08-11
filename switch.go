package main

import "fmt"

func main() {

	name := "Baihaqi"

	switch name {
	case "Yazid":
		fmt.Println("Hello Yazid")
		fmt.Println("Hello Yazid")
	case "John":
		fmt.Println("Hello John")
		fmt.Println("Hello John")
	default:
		fmt.Println("Hi, nice to meet you")
		fmt.Println("Hi, nice to meet you")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Your name is too long")
	case length > 5:
		fmt.Println("middle length")
	default:
		fmt.Println("short length")
	}
}
