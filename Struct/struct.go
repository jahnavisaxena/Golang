package main

import (
	"fmt"
	"time"
)

//orderstruct
type order struct{
	id int
	amount float32
	status stpackage main

import (
	"fmt"
	"time"
)

//orderstruct
type order struct{
	id int
	amount float32
	status string
    createdAt time.Time// nano second precision
}

func newOrder(id int, amount float32,status string)*order{
	myOrder := order{
		id: id,
		amount: amount,
		status: status,	
		
	}

	return &myOrder

}

func (o *order) changeStatus(status string){//receiver type
	o.status = status //struct automatically dereferences
}

func(o order) getAmount() float32{
	return o.amount 
}
func main(){
	
	myOrder:= newOrder(1,90.0,"recieved");
	fmt.Println(myOrder)

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder.status)

	fmt.Println(myOrder.getAmount())
/*
	//myOrder createdAt = time.Now()
	fmt.Println("Order struct", myOrder)
	fmt.Println("id",myOrder.id)

	myOrder2:= order{
		id: 90,
		amount:90.44,
		status:"shipped",
	    createdAt: time.Now(),
	}

	myOrder.status = "paid"

	fmt.Println("Order struct", myOrder2)*/
	
    

}ring
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
