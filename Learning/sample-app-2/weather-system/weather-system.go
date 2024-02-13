// This Program was created by Aravind Nandakumar(2347214) III MCA B

// Imagine you are developing a simple weather application in Go that takes user
// input for the current temperature in Celsius and provides a weather
// recommendation based on the following conditions:
// If the temperature is below 10 degrees Celsius, recommend wearing a heavy
// jacket.
// If the temperature is between 10 and 20 degrees Celsius (inclusive), recommend
// wearing a light jacket.
// If the temperature is above 20 degrees Celsius, recommend wearing a t-shirt.
// Write a Go program that takes the user input for the current temperature,
// processes it using variables and control flow structures, and prints the
// appropriate weather recommendation.
// Your program should include the following:
// Declaration of a variable to store the temperature.
// Input statement to get the temperature from the user.
// Conditional statements to determine the appropriate weather recommendation
// based on the temperature.
// Output statement to display the weather recommendation.

package main

import "fmt"

func main() {
	// choice variable to run the program multiple times
	var choice int32
	choice = 1

	// for loop running until user decides not to
	for choice != 0 {
		fmt.Println("\n-----Weather System-----")
		fmt.Println("Enter the temperature in Celsius:")
		var temp float64
		// taking user input
		fmt.Scan(&temp)

		// control structure to check the temperature
		if temp < 10 {
			fmt.Println("Recommend wearing a heavy jacket.")
		} else if temp >= 10 && temp <= 20 {
			fmt.Println("Recommend wearing a light jacket.")
		} else {
			fmt.Println("Recommend wearing a t-shirt.")
		}

		// asking user to run again or not
		fmt.Println("Do You want to Run again?(1 or 0)")
		fmt.Scan(&choice)
	}
}
