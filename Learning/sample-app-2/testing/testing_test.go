// calculator_test.go

package main

import "testing"

func TestSum(t *testing.T) {
	// Test case 1: Positive numbers
	result := Sum(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("Sum(5, 3) = %d; want %d", result, expected)
	}

	// Test case 2: Negative numbers
	result = Sum(-2, -4)
	expected = -6
	if result != expected {
		t.Errorf("Sum(-2, -4) = %d; want %d", result, expected)
	}

	// Test case 3: Zero values
	result = Sum(0, 0)
	expected = 0
	if result != expected {
		t.Errorf("Sum(0, 0) = %d; want %d", result, expected)
	}
}
