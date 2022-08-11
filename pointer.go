package main

import "fmt"

type Address struct{
	City, Province, Country string
}
func ChangeCountry(address *Address)  {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	// var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa TImur", "Indonesia"}
	
	// fmt.Println(address1)
	// fmt.Println(address2)
	// fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Surabaya"
	address4.Province = "Jawa Timur"
	address4.Country = "Indonesia"
	// fmt.Println(address4)

	address5 := Address {
		City: "Papua",
		Province: "Maluku",
		Country: "",
	}

	ChangeCountry(&address5)
	fmt.Println(address5)
}