package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)

	wg := &sync.WaitGroup{}

	// myCh <- 5

	wg.Add(2)

	// <- recieve only
	go func(myCh <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, isChannelOpen := <-myCh
		// fmt.Println(<-myCh)
		fmt.Println(val, isChannelOpen)
	}(myCh, wg)

	// send only <-
	go func(myCh chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()

		myCh <- 5
		close(myCh)
	}(myCh, wg)

	wg.Wait()

}
