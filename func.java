package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a,b int)int{
	return a-b
}

func getLanguages()(string,string){
	return "go","java"
}
func main() {
	fmt.Println(add(2,5))
	sub(4,7)
}
