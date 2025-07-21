// Channels in Go are built-in data types that allow goroutines to communicate with each other and synchronize execution.

package main

import (
	//"fmt"
	//"time"
	
)

/*
// Ex-1
func recieve(ch chan int){
	i:=5
	ch <- i
}
func main(){
	ch:= make(chan int)
	go recieve(ch)
	res:= <-ch
	fmt.Println(res)
}
*/

/*
//Ex-2
func process(numChan chan int){
	fmt.Println("processign number", <- numChan)
}
func main(){
	numChan:=make(chan int)

	go process(numChan)
	numChan <- 5
	time.Sleep(time.Second*2)
}
*/

/*
// Ex-3 Goroutine Synchronization
func process(done chan bool){
	defer func (){done <- true}()  //we are using defer beause if some sort of error come up in between logic then goroutine will be stuck so to avoid this and func to be fully executed we use defer
	fmt.Println("processing done....")
}

func main(){
	done:= make(chan bool)
	go process(done)
	<-done
}
*/

/*
//Buffer Channel
//  Ex-4
func proc(emailchan chan string){
	
	for email := range emailchan {
		fmt.Println(email)
		time.Sleep(time.Second*2)
	}

}
func main(){
	emailchan:=make(chan string,10)
	for i:=0;i<10;i++{
		emailchan <- fmt.Sprintf("%d@gmail.com",i)
	}
	close(emailchan)
	proc(emailchan)
	
}
*/

/*
//Ex-5
func main(){
	chan1:= make(chan int)
	chan2:= make(chan string)

	go func(){
		chan1 <- 10
	}()
	go func(){
		chan2 <- "hello"
	}()

	for i:=0;i<2;i++{
	select{
	case chan1val:=  <- chan1:
		fmt.Println("accessing channel 1",chan1val)
	
	case chan2val:= <- chan2:
		fmt.Println("accessning channel 2",chan2val)
	}}
}
*/

//Notation for making channel only sending or only recieving
func emailsend(recievee<- chan int, send chan<- int)
			//only can recieve      //only can send 
			