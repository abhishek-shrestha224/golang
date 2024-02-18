package main

import (
	"fmt"
	"time"
)

func main() {

	// !Without the use of go routines => sequential
	greeter("Hello")
	greeter("World")

	fmt.Println("")

	go greeter("Hello from Go Routine")
	go greeter("World from Go Routine")
}

func greeter(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(str)
	}
}
