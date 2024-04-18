package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

// User represents a user of the FitTrack application.
type User struct {
	Username   string
	StepCounts [][]int
	Goals      []int
}

// StorageManager handles the storage and retrieval of user data.
type StorageManager struct {
	users     map[string]User
	saveGroup sync.WaitGroup
}

// NewStorageManager creates a new instance of StorageManager.
func NewStorageManager() *StorageManager {
	return &StorageManager{
		users: make(map[string]User),
	}
}

// RegisterUser registers a new user with the given username.
func (sm *StorageManager) RegisterUser(username string, numDays int, stepGoal int, done chan<- error) {
	// Check if the username is empty
	if username == "" {
		fmt.Println("Username cannot be empty")
		// done <- errors.New("username cannot be empty")
		return
	}

	// Check if numDays or stepGoal is less than or equal to 0
	if numDays <= 0 {
		fmt.Println("Number of days must be greater than 0")
		// done <- errors.New("number of days must be greater than 0")
		return
	}

	if stepGoal < 0 {
		fmt.Println("Step goal cannot be negative")
		// done <- errors.New("step goal cannot be negative")
		return
	}

	// Initialize a new user and add it to the storage
	user := User{
		Username: username,
		Goals:    make([]int, 0),
	}

	// Prompt user for step counts for each day
	for i := 0; i < numDays; i++ {
		var steps int
		fmt.Printf("Enter step count for day %d:", i+1)
		fmt.Scan(&steps)
		if steps < 0 {
			fmt.Println("Steps Count cannot be negative value. Please Enter Again!!")
			i--
			continue
		}
		user.StepCounts = append(user.StepCounts, []int{steps})
	}

	// Set the daily step goal
	user.Goals = append(user.Goals, stepGoal)

	sm.users[username] = user
	done <- nil
}

// GetUserDetails returns the details of a user with the given username.
func (sm *StorageManager) GetUserDetails(username string, done chan<- User, errChan chan<- error) {
	user, exists := sm.users[username]
	if !exists {
		errChan <- errors.New("user not found")
		return
	}
	done <- user
}

// SaveToFile saves the user data to a JSON file.
func (sm *StorageManager) SaveToFile(filename string, done chan<- error) {
	sm.saveGroup.Add(1)
	defer sm.saveGroup.Done()

	// Marshal the user data into JSON format
	data, err := json.MarshalIndent(sm.users, "", "  ")
	if err != nil {
		done <- err
		return
	}

	// Write the data to a JSON file
	file, err := os.Create(filename)
	if err != nil {
		done <- err
		return
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		done <- err
		return
	}

	done <- nil
}

// LoadFromFile loads the user data from a JSON file.
func (sm *StorageManager) LoadFromFile(filename string, done chan<- error) {
	// Open the JSON file for reading
	file, err := os.Open(filename)
	if err != nil {
		// If the file doesn't exist, create an empty map
		sm.users = make(map[string]User)
		done <- nil
		return
	}
	defer file.Close()

	// Decode the JSON data into the users map
	var users map[string]User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		done <- err
		return
	}

	sm.users = users
	done <- nil
}

func main() {
	storage := NewStorageManager()

	// Load user data from file
	loadDone := make(chan error)
	storage.saveGroup.Add(1)
	go func() {
		defer storage.saveGroup.Done()
		storage.LoadFromFile("userdata.json", loadDone)
	}()
	if err := <-loadDone; err != nil {
		fmt.Println("Error loading user data:", err)
	}

	for {
		fmt.Println("\n\n+----FitTrack Menu----+")
		fmt.Println("| 1. New User         |")
		fmt.Println("| 2. See User Details |")
		fmt.Println("| 3. Exit             |")
		fmt.Println("+---------------------+")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var username string
			var numDays, stepGoal int

			fmt.Println("Enter new username:")
			fmt.Scan(&username)
			if _, exists := storage.users[username]; exists {
				fmt.Println("username already exists")
				break
			}

			fmt.Println("Enter number of days to track:")
			fmt.Scan(&numDays)

			fmt.Println("Enter daily step goal:")
			fmt.Scan(&stepGoal)

			registerDone := make(chan error)
			storage.saveGroup.Add(1)
			go func() {
				defer storage.saveGroup.Done()
				storage.RegisterUser(username, numDays, stepGoal, registerDone)
			}()
			err := <-registerDone
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("User registered successfully!")
			}

		case 2:
			var username string
			fmt.Println("Enter username to see details:")
			fmt.Scan(&username)

			getUserDone := make(chan User, 1)
			errChan := make(chan error, 1)
			go storage.GetUserDetails(username, getUserDone, errChan)
			select {
			case user := <-getUserDone:
				fmt.Printf("User Details for %s:\n", username)
				fmt.Println("Step Counts:")
				for i, steps := range user.StepCounts {
					fmt.Printf("Day %d: %d steps\n", i+1, steps[0])
				}
				fmt.Printf("Daily Step Goal: %d\n", user.Goals[0])
			case err := <-errChan:
				fmt.Println("Error:", err)
			}

		case 3:
			// Wait for all pending saves to complete before exiting
			storage.saveGroup.Wait()

			// Save user data to file before exiting
			saveDone := make(chan error)
			go storage.SaveToFile("userdata.json", saveDone)
			if err := <-saveDone; err != nil {
				fmt.Println("Error saving user data:", err)
			}
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
