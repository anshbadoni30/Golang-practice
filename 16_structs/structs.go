package main

import (
	"fmt"
	"time"
)

type customer struct{
	cusid int
	cusname string
} //Using this  for structure  embedding concept

type order struct {
	id int
	status string
	price int
	currentTime time.Time
	customer
}


//If you want to connect your structure to function you have to use "reciever type"
// Take the structure name first letter and full structure name and write after func keyword
// Use * for reeferencing the variable and changing the values of current object variable
func (o *order)changeStatus(status string) {
	o.status=status
}

func(o order) getAmount() int{
	return o.price
}

//Constructor Convention
func newOrder(id int,  status string, price int) *order{
	mything:=order{
		id:id,
		status:status,
		price:price,
	}
	return &mything
}

func main(){
	myOrder:=order {
		id: 1,
		status: "pending",
		price: 200,
		currentTime: time.Now(),
	}
	myOrder2:=order {
		id: 2,
		status: "recieved",
		price: 500,
		currentTime: time.Now(),
	}
	myOrder.changeStatus("cancelled")
	fmt.Println(myOrder2.getAmount())
	fmt.Println(myOrder2, myOrder)

	mything:=newOrder(2,"pending",5000)
	fmt.Println(mything)

	//Alternate method to create and instantiate structure
	language:= struct{
		lname string
		speak bool
	} {"Hindi",true}
	fmt.Println(language)

	//Structure Embedding concept
	mycus1:=customer{
		cusid: 1,
		cusname: "ansh",
	}
	myOrder3:=order{
		id: 3,
		status: "cancelled",
		price: 4567,
		currentTime: time.Now(),
		customer: mycus1,
	}
	fmt.Println(myOrder3)
}