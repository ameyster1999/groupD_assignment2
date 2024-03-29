package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
//ameen function call -500225970
	result := add(3, 7)
	fmt.Println("Sum:", result)

	// Start of Hammad Ul Hassan - 500230292 code
	fmt.Printf("Hammad Ul Hassan - 500230292\n")
	fmt.Printf("Enter a number to check if it's even or odd:\n")
	var num int
	fmt.Scanln(&num)
	result1 := isEven(num)
	fmt.Printf("%d is %s\n", num, result1)
	// End of Hammad Ul Hassan - 500230292 code

	//  start calling sajjad's function
	// Get input from the user
	fmt.Print("Enter temperature in Celsius: ")
	var celsius float64
	fmt.Scanln(&celsius)

	// calling the function
	celsiusToFahrenheit(celsius)
	// calling end for sajjad's function
	// Start of Ramanpreet kaur - 500218959 code
	// Example usage:
	arr := []int{3, 5, 7, 2, 6}

	// Call the sortArray function to sort the array
	sortedArr := sortArray(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", sortedArr)
	// End of Ramanpreet kaur - 500218959 code
	// start of Rishav's code - 500228178
	greatestOfThree()
	// end of Rishav's code - 500228178
	// Start of Sindhuja Peravali - 500228575 code
	fmt.Printf("Sindhuja Peravali - 500228575\n") var num1, num2 int
	// Check for division by zero
	if num2 == 0 {
		fmt.Println("Error: Division by zero is not allowed.") return
	}
	// Call the function and store the return values
	quotient, remainder := divide(num1, num2)
	// Print the results
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder) //End of Sindhuja Peravali - 500228575 code
	// Start of Prateek Nandiwada - 500220142 code
	fmt.Printf("Prateek Nandiwada - 500220142\n")
	var pnum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&pnum)
	if isPrime(pnum) {
		fmt.Printf("%d is a prime number.\n", pnum)
		} else {
		fmt.Printf("%d is not a prime number.\n", pnum)
		}
	// End of Prateek Nandiwada - 500220142 code
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
// ameens code -500225970
// this function is to add 2 number
func add(a, b int) int {
	return a + b
}

// start of KiranJeeth kaur
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

// Created by Ramanpreet kaur - 500218959
// sortArray sorts an array of integers in ascending order
func sortArray(arr []int) []int {
	sort.Ints(arr)
	return arr
}

// End of Ramanpreet kaur - 500218959 code
// rishav code
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

func divide(dividend, divisor int) (quotient, remainder int) {
	// Calculate the quotient
	quotient = dividend / divisor
	// Calculate the remainder
	remainder = dividend % divisor
	return remainder,quotient
}
// Start of Prateek Nandiwada - 500220142 code
// function to find if a number is prime
func isPrime(num int) bool {
if num < 2 {
	return false
	}
limit := int(math.Sqrt(float64(num)))
for i := 2; i <= limit; i++ {
	if num%i == 0 {
	return false
	}
}
return true
}
// End of Prateek Nandiwada - 500220142 code
