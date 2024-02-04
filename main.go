package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	result := add(3, 7)
	fmt.Println("Sum:", result)

	// Start of Hammad Ul Hassan - 500230292 code
	fmt.Printf("Hammad Ul Hassan - 500230292\n")
	fmt.Printf("Enter a number to check if it's even or odd:\n")
	var num int
	fmt.Scanln(&num)
	result := isEven(num)
	fmt.Printf("%d is %s\n", num, result)
	// End of Hammad Ul Hassan - 500230292 code

	//  start calling sajjad's function
	// Get input from the user
	fmt.Print("Enter temperature in Celsius: ")
	var celsius float64
	fmt.Scanln(&celsius)

	// calling the function
	celsiusToFahrenheit(celsius)
	// calling end for sajjad's function

}

// start of sajjad's code
func celsiusToFahrenheit(celsius float64) float64 {
	// conversion from celsius to farenhiet
	fahrenheit := (celsius * 9 / 5) + 32
	// Print the result
	fmt.Printf("%.2f Celsius is equal to %.2f Fahrenheit\n", celsius, fahrenheit)
	return 0
}

// end of sajjad's code

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

// Created by Hammad Ul Hassan - 500230292
// This program checks if a number is even or odd. If the number is even, it returns "Even", else it returns "Odd".
// The logic is simple, if the number is divisible by 2, it is even, else it is odd which is verified by the if statement.
// It asks the user to enter a number and then calls the isEven function to check if the number is even or odd.
// The result is then printed on the screen.
func isEven(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
