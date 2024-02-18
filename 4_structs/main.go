package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (usr User) getName() string {
	return usr.Name
}

func main() {
	abhishek := User{"Abhishek Shrestha", 20, false, "abhi@gmail.com"}

	fmt.Println(abhishek)

}
