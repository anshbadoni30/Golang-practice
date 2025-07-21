package main

import "fmt"

// if you want any datatype in function
/*
func showvalue[T any](value []T){
	for _, num:=range value{
		fmt.Println(num)
	}
}
*/

// if you want certain datatype
func showvalue[T int | string | bool](value []T) {
	for _, num := range value {
		fmt.Println(num)
	}
}

func main() {
	//num:= []int{1,2,3,4}
	//str:=[]string{"Hello","world"}
	bol := []bool{true, false}
	showvalue(bol)
}
