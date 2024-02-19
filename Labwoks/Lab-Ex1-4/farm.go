package main

import (
	"fmt"
)

// Animal struct represents a farm animal with an ID, name, Age, and health.
type Animal struct {
	ID     int
	Name   string
	Age    int
	Health int
	Weight float64
}

// Farm struct represents a farm with a collection of animals.
type Farm struct {
	Animals map[int]Animal
}

// Task represents a task that can be performed on an animal.
type Task func(animal *Animal)

func main() {
	// Initialize the farm
	farm := Farm{
		Animals: make(map[int]Animal),
	}
	fmt.Println("Enter Number of Animals")
	var animalNum int
	fmt.Scan(&animalNum)

	// Take user input to add animals
	fmt.Println("Add Animals to the Farm:")
	for i := 1; i <= animalNum; i++ {
		name := getUserInput(fmt.Sprintf("Enter name for Animal %d: ", i))
		Age := getIntInput(fmt.Sprintf("Enter Age for %s: ", name))
		health := getIntInput(fmt.Sprintf("Enter initial health for(0 - 10) %s: ", name))
		weight := getFloatInput(fmt.Sprintf("Enter Weight of %s", name))

		farm.AddAnimal(i, name, Age, health, weight)
	}

	// Print the initial state of the farm
	fmt.Println("\nInitial Farm State:")
	farm.PrintFarmState()

	// Take user input to perform tasks on animals
	fmt.Println("\nPerform Tasks on Animals:")
	for {
		animalID := getIntInput("Enter Animal ID to perform a task (0 to exit): ")
		if animalID == 0 {
			break
		}

		task := getTaskChoice()
		farm.PerformTask(animalID, task)

		// Print the updated state of the farm
		fmt.Println("\nUpdated Farm State:")
		farm.PrintFarmState()
	}

}

// AddAnimal adds an animal to the farm.
func (f *Farm) AddAnimal(id int, name string, Age int, health int, weight float64) {
	animal := Animal{
		ID:     id,
		Name:   name,
		Age:    Age,
		Health: health,
		Weight: weight,
	}
	f.Animals[id] = animal
}

// PerformTask performs a task on a specific animal.
func (f *Farm) PerformTask(animalID int, task Task) {
	animal, exists := f.Animals[animalID]
	if exists {
		fmt.Printf("\nPerforming task on %s...\n", animal.Name)
		task(&animal)
		f.Animals[animalID] = animal
	} else {
		fmt.Println("\nAnimal not found!")
	}
}

// FeedTask simulates feeding an animal.
func FeedTask(animal *Animal) {
	fmt.Printf("%s is being fed!\n", animal.Name)
	animal.Health += 10
}

// CheckHealthTask checks the health of an animal.
func CheckHealthTask(animal *Animal) {
	fmt.Printf("Checking health of %s...\n", animal.Name)
	if animal.Health < 5 {
		fmt.Printf("%s is not feeling well. Health: %d\n", animal.Name, animal.Health)
	} else {
		fmt.Printf("%s is healthy. Health: %d\n", animal.Name, animal.Health)
	}
}

// ChangeWeightTask increases or decreases the weight of an animal.
func ChangeWeightTask(animal *Animal) {
	fmt.Printf("Current weight of %s: %.2f\n", animal.Name, animal.Weight)

	// Prompt the user to choose an action
	fmt.Println("Choose an action:")
	fmt.Println("1. Increase weight")
	fmt.Println("2. Decrease weight")

	choice := getIntInput("Enter Action Number: ")

	switch choice {
	case 1:
		increaseAmount := getFloatInput("Enter the amount to increase weight: ")
		animal.Weight += increaseAmount
		fmt.Printf("Weight of %s increased by %.2f. New weight: %.2f\n", animal.Name, increaseAmount, animal.Weight)
	case 2:
		decreaseAmount := getFloatInput("Enter the amount to decrease weight: ")
		if decreaseAmount > animal.Weight {
			fmt.Println("Error: Cannot decrease weight below zero.")
			return
		}
		animal.Weight -= decreaseAmount
		fmt.Printf("Weight of %s decreased by %.2f. New weight: %.2f\n", animal.Name, decreaseAmount, animal.Weight)
	default:
		fmt.Println("Invalid choice. No action taken.")
	}
}

// PrintFarmState prints the current state of the farm.
func (f *Farm) PrintFarmState() {
	fmt.Println("Farm Animals:")
	for _, animal := range f.Animals {
		fmt.Printf("[ID: %d | Name: %s | Age: %d | Health: %d | Weight: %f]\n", animal.ID, animal.Name, animal.Age, animal.Health, animal.Weight)
	}
}

// getUserInput takes user input for string values.
func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// getIntInput takes user input for integer values.
func getIntInput(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// getFloatInput takes user input for float values.
func getFloatInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// getTaskChoice gets the user's choice for a task.
func getTaskChoice() Task {
	fmt.Println("Choose a Task:")
	fmt.Println("1. Feed the Animal")
	fmt.Println("2. Check Animal's Health")
	fmt.Println("3. Increase Animal Weight")

	choice := getIntInput("Enter Task Number: ")

	switch choice {
	case 1:
		return FeedTask
	case 2:
		return CheckHealthTask
	case 3:
		return ChangeWeightTask
	default:
		fmt.Println("Invalid choice. Defaulting to FeedTask.")
		return FeedTask
	}
}
