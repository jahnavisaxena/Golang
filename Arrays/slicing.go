package main

import "fmt"

func main() {
	//unitialized slice
	var nums []int

	fmt.Print(nums == nil)
	fmt.Println(len(nums))

	var num = make([]int,2)

	fmt.Println(num )//not nil
}
