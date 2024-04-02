package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJSONMarshaling(t *testing.T) {
	// Create sample farm data
	farm := []Animal{
		{ID: 1, Name: "Cow", Species: "Cow", Color: "Black", Age: 5, IsHealthy: true},
		{ID: 2, Name: "Chicken", Species: "Chicken", Color: "White", Age: 2, IsHealthy: true},
	}

	// Marshal the farm data to JSON
	jsonData, err := json.Marshal(farm)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	// Unmarshal the JSON data back into a slice of animals
	var farmUnmarshaled []Animal
	err = json.Unmarshal(jsonData, &farmUnmarshaled)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	// Compare the original farm data with the unmarshaled farm data
	if !reflect.DeepEqual(farm, farmUnmarshaled) {
		t.Errorf("Expected: %+v, Got: %+v", farm, farmUnmarshaled)
	}
}
