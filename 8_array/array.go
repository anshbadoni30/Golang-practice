package main
import "fmt"

func main(){
	// initializing an array
	var nums[4]int
	//by default in int array the values are zero at all index
	nums[0]=1
	fmt.Println(len(nums)) //printing length of array
	fmt.Println(nums[0])
	fmt.Println(nums)

	var names[4]string
	//by default in string array the values are empty string at all index
	names[2]="ansh"
	fmt.Println(names[2])
	fmt.Println(names)

	//by default in boolean all are false at all index

	//Static Initializatoion
	var arr=[4]int{1,2,3,4}
	fmt.Println(arr)

	//2-D array
	var num1=[2][2]int{{1,2},{3,4}}
	fmt.Println(num1)


}