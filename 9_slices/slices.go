package main

import (
	"fmt"
	"slices"
)

//Slices are just like vector which can increase their size dynamically

func main(){
	//In Slice declaration we dont need to pass size in square brackets
	//uninitialized slice is nil
	var num[]int
	fmt.Println(num) // prints nil

	//General method in golang to make slice
	var nums= make([]int,2,5) //make has 3 parameter that are - datatype, intial length, initial capacity
	//initially nums has 2 numbers which have zero value and has size of 5 to hold
	fmt.Println(nums)
	fmt.Println(cap(nums)) //to see capacity
	fmt.Println(len(nums)) //to see length

	nums=append(nums, 1)
	nums=append(nums, 2)
	nums=append(nums, 3)
	nums=append(nums, 4)
	nums=append(nums, 5)
	fmt.Println(cap(nums))
	fmt.Println(nums)

	//copy function
	var nums2=make([]int,len(nums))
	//nums2=nums easy method
	copy(nums2,nums)
	fmt.Println(nums2)

	//golang also has a package of slice which has multiple functions to perform
	//like append, copy, delete, equal, has, len, cap, make, revers
	fmt.Println(slices.Equal(nums,nums2))
}