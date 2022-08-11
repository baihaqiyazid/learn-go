package main

import "fmt"


func sumAll(numbers ...int) int {
	total := 0
	
	for _, v := range numbers {
		total += v
	}

	return total
}

func main() {
	
	total := sumAll(10,20,30,40,50,60)
	fmt.Println(total)

	numbers := []int {20,30,40,20}
	total = sumAll(numbers...)
	
	fmt.Println(total)
}