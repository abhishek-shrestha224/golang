package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signalList = []string{}

func main() {
	websiteList := []string{
		"https://google.com",
		"https://gsdsdo.dev",
		"https://fb.com",
		"https://github.com"}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // => will stop the main method from terminating until all the go routines are done
}

func getStatusCode(endpoint string) {
	defer wg.Done() // => tells the wait group that this go routine is done with its task and it has no problem with the main method getting terminated

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Something went wrong in %s\n", endpoint)
		return
	}
	mut.Lock() // => prvents anyone else from writing into the memory
	signalList = append(signalList, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
