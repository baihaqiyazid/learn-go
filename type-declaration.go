package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtp NoKTP = "3178293849342"
	var marriedStatus Married = false
	fmt.Println(noKtp)
	fmt.Println(marriedStatus)
}