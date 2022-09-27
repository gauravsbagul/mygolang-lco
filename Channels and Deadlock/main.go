package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Channel in golang")

	myCh := make(chan int, 2) 
	wg := &sync.WaitGroup{}
	
	wg.Add(2)

	//RECEIVED ONLY
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)


		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	//SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 0
		close(myCh)
		// myCh <- 10
		wg.Done()
	}(myCh, wg)

	wg.Wait()


}