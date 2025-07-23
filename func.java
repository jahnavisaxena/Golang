package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a,b int)int{
	return a-b
}

func getLanguages()(string,string,bool){
	return "go","java",true
}
func main() {
	fmt.Println(add(2,5))
	sub(4,7)
	lang1 , lang2 ,truth:=getLanguages()
	fmt.Print(lang1,lang2,truth)
}
