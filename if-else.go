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
}
