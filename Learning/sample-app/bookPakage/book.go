// 3.Case Study: Online Bookstore Inventory Checker
// Problem Statement:
// You're tasked with creating a program for an online bookstore to check the inventory of books.
// The program should allow users to search for books by title and display whether the book is available in the inventory or not.
// Create a struct named Book with title and availability (boolean)
// Create a map named Inventory with Book title as key and availability (boolean) as value
// Check the availability with if condition and print whether the book is available or not

package main

import "fmt"

// Book struct representing a book with title and availability
type Book struct {
	Title        string
	Availability bool
}

func main() {
	// Create an inventory map with Book title as key and availability as value
	inventory := map[string]bool{
		"HarryPotter":   true,
		"Metamorphesis": false,
		"Hamlet":        true,
		"SilentPatient": true,
		"SecretSeven":   false,
		"Goosebumps":    true,
	}

	// Prompt user to enter the title of the book they want to search
	var searchTitle string
	fmt.Print("Enter the title of the book you want to search: ")
	fmt.Scan(&searchTitle)

	// Check if the book is in the inventory
	if availability, ok := inventory[searchTitle]; ok {
		// Book found in the inventory
		if availability {
			fmt.Printf("The book '%s' is available.\n", searchTitle)
		} else {
			fmt.Printf("Sorry, the book '%s' is not available.\n", searchTitle)
		}
	} else {
		// Book not found in the inventory
		fmt.Printf("Sorry, the book '%s' is not found in the inventory.\n", searchTitle)
	}
}
