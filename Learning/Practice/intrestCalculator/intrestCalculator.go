// You want to build a simple interest calculator in GoLang. The program should
// ask the user to input multiple sets of data, each containing the principal amount,
// the annual interest rate, and the number of years for which the interest is to be
// calculated. Use arrays to store the input data for each set, calculate the simple
// interest for each set using the formula: Simple Interest = (Principal Amount *
// Annual Interest Rate * Number of Years) / 100, and display the results.

package main

import (
	"fmt"
)

// Loan struct represents a loan with properties like principal amount, annual interest rate, number of years, and holder name.
type Loan struct {
	PrincipalAmount    float64
	AnnualInterestRate float64
	Years              float64
	HolderName         string
}

func main() {
	var numHolders int
	fmt.Print("Enter the number of loan holders: ")
	fmt.Scan(&numHolders)

	// Array of Loan structs to store input data for each holder
	loans := make([]Loan, numHolders)

	// Take user input for each holder's loan details
	for i := 0; i < numHolders; i++ {
		fmt.Printf("\nEnter loan details for Holder %d:\n", i+1)
		loans[i].HolderName = getStringInput("Holder Name: ")
		loans[i].PrincipalAmount = getUserInput("Principal Amount: ")
		loans[i].AnnualInterestRate = getUserInput("Annual Interest Rate (%): ")
		loans[i].Years = getUserInput("Number of Years: ")
	}

	// Display details for each loan holder
	fmt.Println("\nLoan Details:")
	for i := 0; i < numHolders; i++ {
		displayLoanDetails(loans[i])
	}
}

// getUserInput takes user input for float values.
func getUserInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// getStringInput takes user input for string values.
func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// calculateSimpleInterest calculates the simple interest for a loan.
func calculateSimpleInterest(loan Loan) float64 {
	return (loan.PrincipalAmount * loan.AnnualInterestRate * loan.Years) / 100.0
}

// displayLoanDetails displays the details of a loan holder.
func displayLoanDetails(loan Loan) {
	fmt.Printf("\nHolder: %s\n", loan.HolderName)
	fmt.Printf("Principal Amount: %.2f\n", loan.PrincipalAmount)
	fmt.Printf("Annual Interest Rate: %.2f%%\n", loan.AnnualInterestRate)
	fmt.Printf("Number of Years: %.2f\n", loan.Years)
	fmt.Printf("Simple Interest: %.2f\n", calculateSimpleInterest(loan))
}
