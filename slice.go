package main

import (
	"fmt"
)

func main()  {
	
	var month = [...]string{
		"January",
		"February",
		"March",
		"April",
	}

	fmt.Println(month)
	fmt.Println("length:", len(month))
	fmt.Println("cap:", cap(month))

	var slice1 = month[0:2]
	var appendSlice = append(slice1, "append") //append
	slice1[0] = "newJanuary" //update

	fmt.Println(appendSlice)
	fmt.Println(slice1)

	newSlice := make([]string, len(month), cap(month)) //make Slice
	newSlice[0] = "Monday"
	newSlice[1] = "Tuesday"
	newSlice[2] = "Wednesday"
	newSlice[3] = "Thursday"

	fmt.Println(newSlice)

	newSlice2 := make([]string, len(month), cap(month)) //make Slice
	copy(newSlice2, newSlice) //Copy
	fmt.Println(newSlice2)
}