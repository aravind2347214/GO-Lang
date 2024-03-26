package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID         int
	Name       string
	Email      string
	Hobbies    []string
	Address    Address
	Attributes map[string]string
}

type Address struct {
	City    string
	State   string
	Country string
}

func main() {
	person := Person{
		ID:      1,
		Name:    "Rahul John",
		Email:   "rahul@gamil.com",
		Hobbies: []string{"Reading", "Gardening"},
		Address: Address{
			City:    "New York",
			State:   "NY",
			Country: "USA",
		},
		Attributes: map[string]string{"Age": "30", "Gender": "Male"},
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("Marshalled JSON:")
	fmt.Println(string(jsonData))
}
