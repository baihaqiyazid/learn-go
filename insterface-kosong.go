package main

import "fmt"

func Say(i int) interface{}  {
	if i == 1 {
		return 1
	}else if i == 2{
		return true
	}else {
		return "word"
	}
}

func main()  {
	data := Say(4)
	fmt.Println(data)
}