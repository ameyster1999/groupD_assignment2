package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	result := add(3, 7)
	fmt.Println("Sum:", result)

}
func add(a, b int) int {
	return a + b
}
