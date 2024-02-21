package main

import "fmt"

func calculateVolume(length int) func(int) func(int) int {
	return func(breadth int) func(int) int {
		return func(height int) int {
			return length * breadth * height
		}
	}
}

func closure() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	volume := calculateVolume(4)(5)(6)
	fmt.Println(volume)
	increaseCounter := closure()
	increaseCounter()
	fmt.Println((increaseCounter()))
}
