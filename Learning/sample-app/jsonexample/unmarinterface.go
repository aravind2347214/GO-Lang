package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id         int
	Name       string
	Email      string
	Hobbies    []string
	Address    Address
	Attributes map[string]interface{}
}

type Address struct {
	City    string
	State   string
	Country string
}

func main() {
	jsonData := []byte(`{
		"id": 1,
		"name": "Riya Gupta",
		"email": "riya@gmail.com",
		"hobbies": ["Coding", "Cooking"],
		"address": {
			"city": "San Francisco",
			"state": "Califormia",
			"country": "USA"
		},
		"attributes": {
			"age": 22,
			"height": 175,
			"employed": true
		}
	}`)

	// Unmarshal JSON data into a Person struct
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Accessing the unmarshaled data
	fmt.Println("Unmarshaled Data:")
	fmt.Println("ID:", person.Id)
	fmt.Println("Name:", person.Name)
	fmt.Println("Email:", person.Email)
	fmt.Println("Hobbies:", person.Hobbies)
	fmt.Println("Address:", person.Address)
	fmt.Println("Attributes:", person.Attributes)
}
