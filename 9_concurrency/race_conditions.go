package main

import (
	"fmt"
	"sync"
)

func main() {
	score := []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		mut.Lock()
		score = append(score, 1)
		fmt.Println("GR-1:", score)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		mut.Lock()
		score = append(score, 2)
		fmt.Println("GR-2:", score)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		mut.Lock()
		score = append(score, 3)
		fmt.Println("GR-3:", score)
		mut.Unlock()
	}(wg, mut)

	wg.Wait()
}
