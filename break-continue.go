package main

import "fmt"

func main()  {
	
	for i := 1; i < 10; i++{

		fmt.Println("looping", i)
		if i == 5 {
			fmt.Println("BREAK")
			fmt.Println("")
			break
		}
	}


	for i := 1; i < 10; i++{
		if i%2 != 0 {
			continue
		}
		fmt.Println("looping", i)
	}


}