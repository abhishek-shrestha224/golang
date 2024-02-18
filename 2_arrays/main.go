package main

import (
	"fmt"
	"sort"
)

func slices() {
	var foodSlice = []string{"pizza", "pasta", "sandwich"}
	// fmt.Printf("%T\n", foodSlice)
	fmt.Printf("%d\n", len(foodSlice))
	foodSlice = append(foodSlice, "noodles", "hotdogs", "fries")
	fmt.Printf("%d\n", len(foodSlice))

	fmt.Println(foodSlice[1:4])

	numSlice := make([]int, 5)
	fmt.Println(numSlice, len(numSlice), cap(numSlice))
	numSlice[0] = 9884
	numSlice[1] = 1545
	numSlice[2] = 2121
	numSlice[3] = 8921
	numSlice[4] = 454
	fmt.Println(numSlice)
	fmt.Println(cap(numSlice))
	numSlice = append(numSlice, 123, 4334, 3232, 5656, 76767, 4545, 4343, 545)
	fmt.Println(numSlice, len(numSlice), cap(numSlice))
	sort.Ints(numSlice)

	fmt.Println(numSlice)

	itmToRemove := 2
	numSlice = append(numSlice[:itmToRemove], numSlice[itmToRemove+1:]...)
	fmt.Println(numSlice)
}

func main() {
	// arrNum := [...]int{1, 2, 3, 4, 5}
	// var fruitArr [5]string

	// fruitArr[0] = "pineapple"
	// fruitArr[1] = "banana"
	// fruitArr[2] = "mango"
	// fruitArr[3] = "apple"
	// fruitArr[4] = "peach"

	// fmt.Println(fruitArr)
	// for _, num := range arrNum {
	// 	fmt.Println(num)
	// }

	slices()
	fmt.Println()
}
