package main

import (
	"fmt"
)


func random() interface{} {
	return  9000
}

func main() {
	result := random()
	/* resultToString := result.(string)
	fmt.Println(resultToString) */

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	
	case int:
		fmt.Println("Value", value, "is int")

	case bool:
		fmt.Println("Value", value, "is bool")
	
	default:
		fmt.Println("Unknown Type") 
	
	}
	
}