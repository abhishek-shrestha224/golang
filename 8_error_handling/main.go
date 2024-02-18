package main

import (
	"errors"
	"fmt"
)

func intDivider(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divide by Zero Error")
	}
	return x / y, nil
}

func main() {

	if num, err := intDivider(2, 0); err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println(num)
	}

}
