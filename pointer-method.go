package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Married()  {
	man.Name = "Mr " + man.Name
}

func main() {
	yazid := Man{"Yazid"}
	yazid.Married()

	fmt.Println(yazid.Name)

}