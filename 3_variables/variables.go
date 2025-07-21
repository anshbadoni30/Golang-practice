package main

import "fmt"

func main(){
	//1st method
	var name string= "Ansh Badoni"
	fmt.Println(name)

	//2nd method - taking datatype automatically
	var result = true
	fmt.Println(result)

	//shorthand operators
	nam := "golang"
	num :=54
	res :=false
	fmt.Println(nam,num,res)
	//Note: You can't use shorthand syntax outside the function (globally)
}