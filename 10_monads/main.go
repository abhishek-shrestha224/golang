package main

import (
	"errors"
	"fmt"
)

type Maybe struct {
	value   int
	isValid bool
}

func (this *Maybe) Bind(i int) {
	this.value = i
	this.isValid = true
}

func (this *Maybe) Nothing() {
	this.value = 0
	this.isValid = false
}

func (this Maybe) Map(f func(int) int) (Maybe, error) {
	if this.isValid {
		this.Bind(f(this.value))
		return this, nil
	} else {
		return this, errors.New("Cannot map a function to an empty monad")
	}

}

func main() {
	maybeNum := new(Maybe)
	fmt.Println("Default:", *maybeNum)
	maybeNum.Bind(122)
	fmt.Println("Bind:", *maybeNum)
	// maybeNum.Nothing()
	fmt.Println("Nothin:", *maybeNum)
	doubled, err := maybeNum.Map(func(i int) int { return i * 2 })

	if err != nil {
		fmt.Println(err, doubled)
		return
	}
	fmt.Println(doubled.value)

}
