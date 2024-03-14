package main

import (
	"fmt"
)

// Animal represents a farm animal.
type Animal struct {
	Name   string
	Age    int
	Health string
}

// Function to update animal age using call by value
func updateAgeByValue(animal Animal, newAge int) Animal {
	animal.Age = newAge
	return animal
}

// Function to update animal age using call by reference (pointer)
func updateAgeByReference(animal *Animal, newAge int) {
	animal.Age = newAge
}

// Function to update animal health using call by value
func updateHealthByValue(animal Animal, newHealth string) Animal {
	animal.Health = newHealth
	return animal
}

// Function to update animal health using call by reference (pointer)
func updateHealthByReference(animal *Animal, newHealth string) {
	animal.Health = newHealth
}

// Function to print animal details using call by value
func printAnimalDetailsByValue(animal Animal) {
	fmt.Printf("Name: %s, Age: %d, Health: %s\n", animal.Name, animal.Age, animal.Health)
}

// Function to print animal details using call by reference (pointer)
func printAnimalDetailsByReference(animal *Animal) {
	fmt.Printf("Name: %s, Age: %d, Health: %s\n", animal.Name, animal.Age, animal.Health)
}

func main() {
	var numAnimals int

	// Get the number of animals from the user
	fmt.Print("Enter the number of animals: ")
	fmt.Scan(&numAnimals)

	// Create an array to store the animals
	animals := make([]Animal, numAnimals)

	// Allow the user to enter details for each animal
	for i := 0; i < numAnimals; i++ {
		var name string
		var age int
		var health string

		// Get details from the user
		fmt.Printf("\nEnter details for Animal %d:\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Age: ")
		fmt.Scan(&age)
		fmt.Print("Health: ")
		fmt.Scan(&health)

		// Initialize the current animal in the array
		animals[i] = Animal{Name: name, Age: age, Health: health}
	}

	// Display a menu for the user to choose operations
	for {
		fmt.Println("\nChoose an operation:")
		fmt.Println("1. Print details for all animals")
		fmt.Println("2. Update age for an animal (Call by Value)")
		fmt.Println("3. Update age for an animal (Call by Reference)")
		fmt.Println("4. Update health for an animal (Call by Value)")
		fmt.Println("5. Update health for an animal (Call by Reference)")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Print details for all animals
			fmt.Println("\nDetails of Entered Animals:")
			for i, animal := range animals {
				fmt.Printf("Animal %d:\n", i+1)
				printAnimalDetailsByValue(animal)
			}
		case 2:
			// Update age for an animal (Call by Value)
			var animalIndex int
			var newAge int

			fmt.Print("\nEnter the index of the animal to update age: ")
			fmt.Scan(&animalIndex)
			fmt.Print("Enter the new age: ")
			fmt.Scan(&newAge)

			if animalIndex >= 1 && animalIndex <= numAnimals {
				animals[animalIndex-1] = updateAgeByValue(animals[animalIndex-1], newAge)
				fmt.Println("Age updated successfully! (Call by Value)")
			} else {
				fmt.Println("Invalid index.")
			}
		case 3:
			// Update age for an animal (Call by Reference)
			var animalIndex int
			var newAge int

			fmt.Print("\nEnter the index of the animal to update age: ")
			fmt.Scan(&animalIndex)
			fmt.Print("Enter the new age: ")
			fmt.Scan(&newAge)

			if animalIndex >= 1 && animalIndex <= numAnimals {
				updateAgeByReference(&animals[animalIndex-1], newAge)
				fmt.Println("Age updated successfully! (Call by Reference)")
			} else {
				fmt.Println("Invalid index.")
			}
		case 4:
			// Update health for an animal (Call by Value)
			var animalIndex int
			var newHealth string

			fmt.Print("\nEnter the index of the animal to update health: ")
			fmt.Scan(&animalIndex)
			fmt.Print("Enter the new health: ")
			fmt.Scan(&newHealth)

			if animalIndex >= 1 && animalIndex <= numAnimals {
				animals[animalIndex-1] = updateHealthByValue(animals[animalIndex-1], newHealth)
				fmt.Println("Health updated successfully! (Call by Value)")
			} else {
				fmt.Println("Invalid index.")
			}
		case 5:
			// Update health for an animal (Call by Reference)
			var animalIndex int
			var newHealth string

			fmt.Print("\nEnter the index of the animal to update health: ")
			fmt.Scan(&animalIndex)
			fmt.Print("Enter the new health: ")
			fmt.Scan(&newHealth)

			if animalIndex >= 1 && animalIndex <= numAnimals {
				updateHealthByReference(&animals[animalIndex-1], newHealth)
				fmt.Println("Health updated successfully! (Call by Reference)")
			} else {
				fmt.Println("Invalid index.")
			}
		case 6:
			// Exit the program
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
