package main

import "fmt"

func outer() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

func main() {
    counter := outer()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
