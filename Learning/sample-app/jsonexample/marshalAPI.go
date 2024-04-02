package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Joke represents a single joke retrieved from the API.
type Joke struct {
	ID        int
	Type      string
	Setup     string
	Punchline string
}

func main() {
	// Fetch jokes from the API
	jokes, err := fetchJokes()
	if err != nil {
		fmt.Println("Error fetching jokes:", err)
		return
	}

	// Display fetched jokes
	fmt.Println("Fetched Jokes:")
	for _, joke := range jokes {
		fmt.Printf("ID: %d\nType: %s\nSetup: %s\nPunchline: %s\n\n", joke.ID, joke.Type, joke.Setup, joke.Punchline)
	}

	// Marshal jokes to JSON
	jsonData, err := json.Marshal(jokes)
	if err != nil {
		fmt.Println("Error marshaling jokes to JSON:", err)
		return
	}

	fmt.Println("JSON representation of jokes:")
	fmt.Println(string(jsonData))

	// Unmarshal JSON data back into jokes
	var unmarshaledJokes []Joke
	err = json.Unmarshal(jsonData, &unmarshaledJokes)
	if err != nil {
		fmt.Println("Error unmarshaling JSON to jokes:", err)
		return
	}

	fmt.Println("\nUnmarshaled Jokes:")
	for _, joke := range unmarshaledJokes {
		fmt.Printf("ID: %d\nType: %s\nSetup: %s\nPunchline: %s\n\n", joke.ID, joke.Type, joke.Setup, joke.Punchline)
	}
}

// Fetch jokes from the API
func fetchJokes() ([]Joke, error) {
	url := "https://official-joke-api.appspot.com/jokes/programming/random"

	// Make a GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch jokes from API: %v", err)
	}
	defer resp.Body.Close()

	// Decode JSON response
	var jokes []Joke
	err = json.NewDecoder(resp.Body).Decode(&jokes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}
	return jokes, nil
}
