package main

import (
	"fmt"
	"maps"
)

func main(){
	var m=make(map[string]string)
	m["name"]="Ansh"
	m["lname"]="badoni"
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["phone"]) //getting a value of key which is not in map
	//will get nil in case of string, 0 in case of int and false in boolean case

	fmt.Println(len(m)) //length of map
	//delete(m,"lname") //deletes the specific item from map
	fmt.Println(m)
	//clear(m) //emty the map
	fmt.Println(m)

	m1:=make(map[string]string)
	m1["name"]="Ansh"
	m1["lname"]="badoni"
	fmt.Println(maps.Equal(m,m1))
}