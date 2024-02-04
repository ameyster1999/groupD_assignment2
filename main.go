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

// Start of Sindhuja Peravali - 500228575 code
// function to calculator quotient and remainder of two numbers 
func divide(dividend, divisor int) (quotient, remainder int) {
	// Calculate the quotient 
	quotient = dividend / divisor
	// Calculate the remainder 
	remainder = dividend % divisor return
}
// End of Sindhuja Peravali - 500228575 code
