package main
import "fmt"

func main(){
	//Constants once assigned, can't be assigned again
	const age = 30
	// age=40 will give error
	fmt.Println(age)

	//constant grouping
	const (
		firstName = "John"
		lastName = "Doe"
	)
	fmt.Println(firstName, lastName)
}