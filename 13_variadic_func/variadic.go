package main
import "fmt"

func sum(nums ...int) int{
	total:=0
	for _,num:=range nums {
		total+=num
	}
	return total
}


func main(){
	result:=sum(1,2,3,4,5,6)
	fmt.Println(result)

	var numi=[]int{22,33}
	result1:=sum(numi...)
	fmt.Println(result1)
}