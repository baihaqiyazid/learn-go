package main

import "fmt"

func main()  {

	firstName := "Yazid"
	lastName := "Baihaqi"
	var age int8 = 20

	firstLastName := string(lastName[0])
	conv_Age := int16(age) 
	
	fmt.Println(firstName, firstLastName,conv_Age)

}