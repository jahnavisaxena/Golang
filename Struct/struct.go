package main

import (
	"fmt"
	"time"
)

//orderstruct
type order struct{
	id int
	amount float32
	status string
	cretedAt time.Time// nano second precision
}
func main(){
	myOrder := order{
		id: 90,
		amount: 34.5,
		status: "received",
		
	}
	fmt.Println("Order struct", myOrder)
}
