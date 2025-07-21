package main

import (
	"fmt"
	"sync"
)

type email struct{
	views int
	mu sync.Mutex
}

func (e *email) increase(wg *sync.WaitGroup){
	defer func(){
		e.mu.Unlock()
		wg.Done()
	}()
	e.mu.Lock()
	e.views+=1
	
}

func main(){
	myemail:=email{
		views: 0,
	}
	var wg sync.WaitGroup
	for i:=0;i<100;i++{
		wg.Add(1)
		go myemail.increase(&wg)
	}
	wg.Wait()
	fmt.Println(myemail.views)
}