package main
import "fmt"

func main(){
	// there is only for loop for looping
	//we can contruct while using for loop

	//While loop
	i:= 0
	for i<=5{
		fmt.Println(i)
		i++
	}

	//infinte loop
	/*for {
		fmt.Println("infinite loop")
	}*/

	// for loop
	for j:=0;j<=3;j++{
		fmt.Println(j)
	}

	//using range with for loop
	for k:= range 10{
		fmt.Println(k)
	}
}