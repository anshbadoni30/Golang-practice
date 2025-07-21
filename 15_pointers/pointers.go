package main
import "fmt"

func sum(nums *int){
	*nums=5
	fmt.Println("the value of nums is ",*nums)
}

func main(){
	numi:=1
	fmt.Println("the value of numi is ",numi)
	sum(&numi)
	fmt.Println("the value of numi is ",numi)
}