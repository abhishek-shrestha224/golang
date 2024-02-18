package main

import "fmt"

const (
	_ = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	today := Sunday

	fmt.Println(today)
}
