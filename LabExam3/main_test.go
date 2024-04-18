package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	storage := NewStorageManager()

	// Test case: Registering a new user with valid input
	errChan := make(chan error)
	go storage.RegisterUser("testuser", 3, 100, errChan)
	err := <-errChan
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	// Test case: Registering a new user with negative step goal
	go storage.RegisterUser("user1", 3, -100, errChan)
	err = <-errChan
	if err == nil {
		t.Error("Expected error for registering user with negative step goal, but got nil")
	}

	// Test case: Registering a new user with zero numDays
	go storage.RegisterUser("user2", 0, 5000, errChan)
	err = <-errChan
	if err == nil {
		t.Error("Expected error for registering user with zero numDays, but got nil")
	}

	// Test case: Registering a new user with empty username
	go storage.RegisterUser("", 3, 5000, errChan)
	err = <-errChan
	if err == nil {
		t.Error("Expected error for registering user with empty username, but got nil")
	}
}

func TestSerializationDeserialization(t *testing.T) {
	storage := NewStorageManager()

	// Register a user with empty name
	errChan := make(chan error)
	go storage.RegisterUser("", 10, 100, errChan)
	<-errChan

	// Save user data to file
	saveDone := make(chan error)
	go storage.SaveToFile("test_userdata.json", saveDone)
	err := <-saveDone
	if err != nil {
		t.Errorf("Failed to save user data: %v", err)
	}

	// Load user data from file
	loadDone := make(chan error)
	go storage.LoadFromFile("test_userdata.json", loadDone)
	err = <-loadDone
	if err != nil {
		t.Errorf("Failed to load user data: %v", err)
	}

	// Compare loaded data with original data
	loadedUser := storage.users["testuser"]
	if len(loadedUser.StepCounts) != 3 || loadedUser.Goals[0] != 10000 {
		t.Errorf("Loaded user data does not match original data")
	}
}
