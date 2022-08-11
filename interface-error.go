package main

import (
	"errors"
	"fmt"
)

func Divide(value int, div int) (int, error)  {
	if div == 0 {
		return 0, errors.New("Divide no zero")
	}else {
		result := value / div
		return result, nil
	}
}

func main() {
	
	result, err := Divide(10, 0)
	if err == nil{
		fmt.Println("Result", result)
	}else {
		fmt.Println("Error", err.Error())
	}
}