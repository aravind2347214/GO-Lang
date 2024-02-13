// This program is Made By Aravind Nandakumar (2347214)
// 2.Use Switch case for branching
// Create User defined functions for every conversion ( 6 User Defined Function)
// Case Study: Scientific Calculator
// Problem Statement:
// Develop a Scientific calculator program in Go that performs

// 1. Addition (+)
// 2. Subtraction (-)
// 3. Multiplication (*)
// 4. Division (/)
// 5. Exponentiation (x^y)
// 6. Square root (√x)
package main

import (
	"fmt"
	"math"
)

// Function to perform addition
func add(x, y float64) float64 {
	return x + y
}

// Function to perform subtraction
func subtract(x, y float64) float64 {
	return x - y
}

// Function to perform multiplication
func multiply(x, y float64) float64 {
	return x * y
}

// Function to perform division
func divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	fmt.Println("Error: Division by zero")
	return 0
}

// Function to perform exponentiation
func exponentiate(x, y float64) float64 {
	return math.Pow(x, y)
}

// Function to calculate square root
func squareRoot(x float64) float64 {
	if x >= 0 {
		return math.Sqrt(x)
	}
	fmt.Println("Error: Cannot calculate square root of a negative number")
	return 0
}

// Main function where the program starts execution
func main() {
	var choice int
	var num1, num2 float64

	// Display the menu
	fmt.Println("Scientific Calculator")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exponentiation (x^y)")
	fmt.Println("6. Square root (√x)")

	// Prompt user for choice and input numbers
	fmt.Print("Enter your choice (1-6): ")
	fmt.Scan(&choice)

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	// Check if the chosen operation requires a second number
	if choice != 6 {
		fmt.Print("Enter the second number: ")
		fmt.Scan(&num2)
	}

	// Switch statement to perform the selected operation
	switch choice {
	case 1:
		result := add(num1, num2)
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
	case 2:
		result := subtract(num1, num2)
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case 3:
		result := multiply(num1, num2)
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case 4:
		result := divide(num1, num2)
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
	case 5:
		result := exponentiate(num1, num2)
		fmt.Printf("%.2f ^ %.2f = %.2f\n", num1, num2, result)
	case 6:
		result := squareRoot(num1)
		fmt.Printf("√%.2f = %.2f\n", num1, result)
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}
