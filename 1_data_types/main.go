package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func runes() {
	myStr := "Abhishek Shrestha"
	// fmt.Println(len(myStr))
	fmt.Println(utf8.RuneCountInString(myStr))
	for i, runeVal := range myStr {
		fmt.Printf("%d\t: %#U \n", i, runeVal)
	}

}
func main() {

	var firstName string = "Abhishek"
	lastName := "Shrestha"

	age, isLoggedIn := 20, false //multi var initilization

	pi := 3.14
	num := int(pi)
	num2, err := strconv.Atoi(lastName)
	fmt.Println(firstName, lastName)
	fmt.Println(isLoggedIn, age)
	fmt.Println(int64(num))
	fmt.Println(num2, err)

	runes()
}
