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
	//  result := isEven(num)
	fmt.Printf("%d is %s\n", num, isEven(num))
	// End of Hammad Ul Hassan - 500230292 code

	print("\n*** Rishav's code ***\n")

	// start of Rishav's code - 500228178
	greatestOfThree()
	// end of Rishav's code - 500228178
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

func greatestOfThree() {
	// Prompt user for input
	fmt.Println("Enter three numbers:")

	// Taking three variables of integer type
	var num1, num2, num3 int

	// Read inputs from user and save in num1, num2 and num3 variables
	fmt.Print("Number 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Number 2: ")
	fmt.Scanln(&num2)
	fmt.Print("Number 3: ")
	fmt.Scanln(&num3)

	// Calling the 'greatest' function and storing the return value in num variable
	num := greatest(num1, num2, num3)

	// Displaying the output on the screen
	fmt.Println(num, "is the greatest\n")
}

// function to calculate the greatest out of three
func greatest(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
