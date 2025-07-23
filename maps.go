package main

import (
	"fmt"
	"maps"
)

//maps ->associative data structure ,hash,objects,dictionary

func main(){
	m := make(map[string]string)//[key]value

	m["name"] ="golang"
	m["area"] = "backend"

	delete(m,"name")

	clear(m)
	//fmt.Println(m["name"],m["area"])
	//fmt.Println(m["phone"])

	//m1:=make(map[string]int)
	//m1["Age"] = 30

	//fmt.Print(m)
	//fmt.Print(len(m))

	a := map[string]int{"price":50, "age":59}
	a1:= map[string]int{"price":50, "age":59}

	fmt.Println(maps.Equal(a,a1))
	k,ok := a["price"]

	fmt.Println(k)

	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}
	
}                          
