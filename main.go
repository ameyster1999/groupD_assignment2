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
func revString(str string) string {
	var theArray []string
	theArray = strings.Split(str, "")
	var strOutput string
	for i := len(str) - 1; i >= 0; i-- {
		strOutput += theArray[i]
	}
	return strOutput
}
