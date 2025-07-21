package main

import "fmt"

//Normal function - Type-1
func add(a int, b int) int {
	return a + b + a
}

//function which returns multiple datatype values - Type-2
func summ(a int, b int) (int, string, bool) {
	return a + b, "hello world", true
}

// Define a function that takes another function as a parameter - Type-3
func applyFunction(x int, fn func(int) int) int {
	return fn(x)
}

// A sample function to be passed
func square(n int) int {
	return n * n
}

//function which return another function - Type-4
func outerFunction() func(int) int {
	return func(x int) int {
		return x * x
	}
}

func main() {
	ans1 := add(3, 4)
	fmt.Println(ans1)
	a1, a2, a3 := summ(4, 5)
	fmt.Println(a1, a2, a3)

	fmt.Println(applyFunction(5, square))

	squareFunc := outerFunction() // outerFunction returns a function
	result := squareFunc(5)       // call the returned function
	fmt.Println("Result:", result)
}
