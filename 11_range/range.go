package main
import "fmt"

func main(){

	var nums= []int{1,2,3,4}
	for i:=0;i<len(nums);i++{
		fmt.Println(nums[i])
	}
	//Use of range function on slice
	for a,b:=range nums{
		fmt.Println(a,b) //a give index and b give value
	}

	//use on maps
	m:= map[string]string{"name":"Ansh", "lname":"badoni"}
	for key,value:=range m{
		fmt.Println(key,value) //in case of maps key and value are written from range
		//In case of single variable, key is written
	}

	//use on strings
	str:= "Hello"
	for i,j:= range str{
		fmt.Println(i,j)  //in case of string first var return index and 2nd var return unicode
	}
}