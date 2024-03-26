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
	jsonData := []byte(`{
		"ID": 1,
		"Name": "Vanessa Mathews",
		"Email": "vanessa@gmail.com",
		"Hobbies": ["Films", "Painting"],
		"Address": {
			"City": "Tulsa",
			"State": "Okhlahoma",
			"Country": "USA"
		},
		"Attributes": {"Age": "25", "Gender": "Female"}
	}`)

	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Unmarshaled Struct:", person)
}
