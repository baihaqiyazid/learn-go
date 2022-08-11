package main

import "fmt"

func endApp()  {
	message := recover() //get panic from App
	if message != nil {
		fmt.Println(message)
	}
	fmt.Println("End App")
}

func App(error bool)  {
	defer endApp() //Still run through error
	if error {
		panic("Something Error")
	}
	fmt.Println("run App")	
}

func main() {
	App(true)
}