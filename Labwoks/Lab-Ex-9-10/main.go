package main

import (
	"fmt"
	"sync"
)

type Animal struct {
	Name    string
	Species string
	Age     int
	Health  string
	Weight  float64
}

type Farm struct {
	Animals []*Animal
}

func (farm *Farm) AddAnimal(animal *Animal) {
	farm.Animals = append(farm.Animals, animal)
}

func feedAnimal(wg *sync.WaitGroup, animal *Animal, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("Feeding %s the %s", animal.Name, animal.Species)
}

func main() {
	farm := &Farm{}

	// Take user input for the number of animals
	var numAnimals int
	fmt.Print("Enter the number of animals: ")
	fmt.Scan(&numAnimals)

	// Adding animals to the farm based on user input
	for i := 0; i < numAnimals; i++ {
		var name, species, health string
		var age int
		var weight float64

		fmt.Printf("\nEnter details for animal %d:\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Species: ")
		fmt.Scan(&species)
		fmt.Print("Age: ")
		fmt.Scan(&age)
		fmt.Print("Health: ")
		fmt.Scan(&health)
		fmt.Print("Weight: ")
		fmt.Scan(&weight)

		animal := &Animal{Name: name, Species: species, Age: age, Health: health, Weight: weight}
		farm.AddAnimal(animal)
	}

	// Create a channel to communicate feeding tasks

	fmt.Println("All Animals are Fed Together and the Order is Based On System Speed")
	ch := make(chan string)

	var wg sync.WaitGroup

	// Start Goroutines to feed animals concurrently
	for _, animal := range farm.Animals {
		wg.Add(1)
		go feedAnimal(&wg, animal, ch)
	}

	// Close the channel when all feeding tasks are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive and print messages from the channel
	for message := range ch {
		fmt.Println(message)
	}
}
