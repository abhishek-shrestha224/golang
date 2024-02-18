package main

import "fmt"

func main() {
	capitalCityHash := make(map[string]string)

	capitalCityHash["Nepal"] = "Kathmandu"
	capitalCityHash["China"] = "Beijing"
	capitalCityHash["France"] = "Paris"

	fmt.Println(capitalCityHash)
	delete(capitalCityHash, "France")
	fmt.Println(capitalCityHash)

	for country, capital := range capitalCityHash {
		fmt.Println(country, ":", capital)
	}
}
