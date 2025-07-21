package main

import "fmt"

func main() {
	age := 8
	if age >= 18 {
		fmt.Print("an adult")
	}else if age >= 12{
		fmt.Print("teenager")
	}else{
		fmt.Print("kid")
	}
	var role ="admin"

	var hasPermission =false

	if role == "admin" && hasPermission{
		fmt.Print("yes")
	}

	if age:= 15; age >= 18 {
		fmt.Print("adult",age)
	}else if age >=12 {
		fmt.Print("teen",age)
	}
}
