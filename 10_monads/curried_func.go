package main

import "fmt"

func calculateVolume(length int) func(int) func(int) int {
	return func(breadth int) func(int) int {
		return func(height int) int {
			return length * breadth * height
		}
	}
}
func main() {
	// Example usage
	volume := calculateVolume(4)(5)(6)
	fmt.Println(volume) // Output: 120
}
