package main

import (
	"fmt"
)

// Animal represents a farm animal.
type Animal struct {
	Name    string
	Species string
	Age     int
	Health  string
	Weight  float64
}

// FarmTask is an interface for performing tasks on animals.
type FarmTask interface {
	PerformTask(animal *Animal) error
}

// FeedTask represents the feeding task.
type FeedTask struct{}

// PerformTask feeds the animal.
func (f *FeedTask) PerformTask(animal *Animal) error {
	fmt.Printf("Feeding %s the %s.\n", animal.Name, animal.Species)
	return nil
}

// MakeSoundTask represents the make sound task.
type MakeSoundTask struct{}

// PerformTask makes the animal produce a sound.
func (m *MakeSoundTask) PerformTask(animal *Animal) error {
	fmt.Printf("%s the %s is making a sound!\n", animal.Name, animal.Species)
	return nil
}

// CheckHealthTask represents the check health task.
type CheckHealthTask struct{}

// PerformTask checks the health of the animal.
func (c *CheckHealthTask) PerformTask(animal *Animal) error {
	fmt.Printf("Checking health of %s the %s. Health: %s\n", animal.Name, animal.Species, animal.Health)
	return nil
}

// ChangeWeightTask represents the change weight task.
type ChangeWeightTask struct{}

// PerformTask changes the weight of the animal.
func (cw *ChangeWeightTask) PerformTask(animal *Animal) error {
	fmt.Printf("Changing weight of %s the %s. Current Weight: %.2f\n", animal.Name, animal.Species, animal.Weight)
	return nil
}

// TaskManager manages farm tasks.
type TaskManager struct{}

// PerformTask executes a task on the given animal.
func (t *TaskManager) PerformTask(task FarmTask, animal *Animal) error {
	return task.PerformTask(animal)
}

// FarmManager manages the farm and performs tasks.
type FarmManager struct {
	Animals []*Animal
}

// AddAnimal adds an animal to the farm.
func (f *FarmManager) AddAnimal(name, species string, age int, health string, weight float64) {
	animal := &Animal{
		Name:    name,
		Species: species,
		Age:     age,
		Health:  health,
		Weight:  weight,
	}
	f.Animals = append(f.Animals, animal)
}

// GetAnimal retrieves an animal by name.
func (f *FarmManager) GetAnimal(name string) (*Animal, error) {
	for _, animal := range f.Animals {
		if animal.Name == name {
			return animal, nil
		}
	}
	return nil, &AnimalNotFoundError{AnimalName: name}
}

// PerformTaskOnFarm performs a task on an animal in the farm.
func (f *FarmManager) PerformTaskOnFarm(task FarmTask, animalName string) error {
	animal, err := f.GetAnimal(animalName)
	if err != nil {
		return err
	}

	return task.PerformTask(animal)
}

// AnimalNotFoundError represents an error when an animal is not found.
type AnimalNotFoundError struct {
	AnimalName string
}

// Error returns the error message for AnimalNotFoundError.
func (e *AnimalNotFoundError) Error() string {
	return fmt.Sprintf("Animal not found: %s", e.AnimalName)
}

func main() {
	// Create a farm manager.
	farmManager := &FarmManager{}

	// Take user input for the number of animals.
	var numAnimals int
	fmt.Print("Enter the number of animals: ")
	fmt.Scan(&numAnimals)

	// Take user input for adding animals.
	for i := 0; i < numAnimals; i++ {
		var name, species, health string
		var age int
		var weight float64

		fmt.Printf("\nEnter details for animal %d:\n", i+1)
		fmt.Print("Enter name: ")
		fmt.Scan(&name)
		fmt.Print("Enter species: ")
		fmt.Scan(&species)
		fmt.Print("Enter age: ")
		fmt.Scan(&age)
		fmt.Print("Enter health status: ")
		fmt.Scan(&health)
		fmt.Print("Enter weight: ")
		fmt.Scan(&weight)

		farmManager.AddAnimal(name, species, age, health, weight)
	}

	// Take user input for tasks.
	for {
		var taskType, animalName string
		fmt.Print("\nEnter task type (Feed/MakeSound/CheckHealth/ChangeWeight/Exit): ")
		fmt.Scan(&taskType)

		if taskType == "Exit" {
			break
		}

		fmt.Print("Enter animal name: ")
		fmt.Scan(&animalName)

		var task FarmTask

		switch taskType {
		case "Feed":
			task = &FeedTask{}
		case "MakeSound":
			task = &MakeSoundTask{}
		case "CheckHealth":
			task = &CheckHealthTask{}
		case "ChangeWeight":
			task = &ChangeWeightTask{}
		default:
			fmt.Println("Invalid task type.")
			continue
		}

		// Perform tasks on animals.
		err := farmManager.PerformTaskOnFarm(task, animalName)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
