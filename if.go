package main

import (
	"fmt"
)

func main() {
	var name = "Baihaqi"

	if name == "Yazid" {
		fmt.Println("Hello Yazid")
	} else if name == "John" {
		fmt.Println("Hello John")
	} else if name == "Tom" {
		fmt.Println("Hello Tom")
	} else {
		fmt.Println("Hi,", name)
	}

	const x = 5
	if length  := len(name); length < x {
		fmt.Println("Length of your name >", x)
	}else {
		fmt.Println("Length of your name is", length)
	}

}