package main

import (
	"encoding/json"
	"fmt"
)

// Animal represents an animal in the farming domain.
type Animal struct {
	ID        int
	Name      string
	Species   string
	Color     string
	Age       int
	IsHealthy bool
}

func main() {
	// Initialize a slice to store animal data
	var Farm []Animal

	// Prompt the user to enter details for multiple animals
	for i := 1; ; i++ {
		fmt.Printf("Enter details for Animal %d:\n", i)
		animal := Animal{}

		// Get user input for animal details
		fmt.Printf("ID: ")
		fmt.Scan(&animal.ID)
		fmt.Printf("Name: ")
		fmt.Scan(&animal.Name)
		fmt.Printf("Species: ")
		fmt.Scan(&animal.Species)
		fmt.Printf("Color: ")
		fmt.Scan(&animal.Color)
		fmt.Printf("Age: ")
		fmt.Scan(&animal.Age)
		fmt.Printf("Is the animal healthy? (true/false): ")
		fmt.Scan(&animal.IsHealthy)

		// Add the animal to the farm
		Farm = append(Farm, animal)

		// Ask if the user wants to add another animal
		var another string
		fmt.Printf("Add another animal? (yes/no): ")
		fmt.Scan(&another)
		if another != "yes" {
			break
		}
	}

	// Marshal the Farm slice to JSON
	jsonData, err := json.Marshal(Farm)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println("JSON Data (Marshalled):")
	fmt.Println(string(jsonData))

	// Unmarshal the JSON data back into a slice of animals
	var FarmUnmarshaled []Animal
	err = json.Unmarshal(jsonData, &FarmUnmarshaled)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the unmarshaled data
	fmt.Println("\nUnmarshaled Data:")
	for _, animal := range FarmUnmarshaled {
		fmt.Printf("ID: %d, Name: %s, Species: %s, Color: %s, Age: %d, IsHealthy: %t\n", animal.ID, animal.Name, animal.Species, animal.Color, animal.Age, animal.IsHealthy)
	}
}
