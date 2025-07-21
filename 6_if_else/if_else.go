package main

import "fmt"

func main(){
	age:=2

	if age<20 {
		fmt.Println("You are a teenager")
	} else if age==20 {
		fmt.Println("teenage ends")
	} else {
		fmt.Println("You are an adult")
	}

	//We can declare a variabel in if construct
	if age1:=21; age1>=12 && age1<=19{
		fmt.Println("Person is teenager")
	}else if age1>=20{
		fmt.Println("Person is an adult")
	}
	
	//Go doesnt have ternary operator
}