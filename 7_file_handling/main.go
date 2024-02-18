package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ifErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	ifErrPanic(err)

	fmt.Println(string(databyte))
}

func main() {
	content := "My name is Abhishek Shrestha"

	file, err := os.Create("./my_file.txt")
	ifErrPanic(err)

	length, err := io.WriteString(file, content)
	ifErrPanic(err)

	fmt.Println(length)
	readFile("./my_file.txt")

	defer file.Close()
}
