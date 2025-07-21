package main

import (
	"fmt"
	"time"
)

func main(){
	i:=2
	// Normal switch
	switch i{
	case 2:
		fmt.Println("i is 2")
	case 1:
		fmt.Println("i is 1")
	default:
		fmt.Println("i can be anything")
	}

	//Multiple conditions on switch
	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend!"	)
		default:
			fmt.Println("It's a weekday.")
	}
	// type condition
	k := func(ii interface{}){
	switch ii.(type){
		case int:
			fmt.Println("k is int")
		case bool:
			fmt.Println("k is bool")
		default:
			fmt.Println("k can any datatype")
	}}
	k(true)
}