package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name){
		fmt.Println(name, "Blocked")
	}else {
		fmt.Println("Welcome", name)
	}	
}

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("Yazid", blackList)
	registerUser("admin", blackList)
}