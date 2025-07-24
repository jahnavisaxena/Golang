package main

import "fmt"

func sum(nums ...int)int{
	total :=0

	for _,num := range nums{
		total+=num

	}
	return total
}
func main() {
	//fmt.Println(1,2,3,4,5,6,7,8,9,0)
	res:=sum(12,4,45,2,323,4,5,5)
	fmt.Println(res)
}
