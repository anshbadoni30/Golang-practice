package main

import "fmt"

/*type current int

const(
	Recieved current= iota
	Pending
	Rejected	
	cancelled
)
*/

type current string
const(
	Recieved current = "Recieved"
	Pending current = "Pending"
	Rejected current = "Rejected"
)

func orderstatus(status current){
	fmt.Println("your order status is",status)
}

func main(){

	orderstatus(Rejected)
}