// Case Study: Student Grade Calculator

// Problem Statement:
// You're tasked with creating a program to calculate the final grades of students based on their scores in multiple subjects.
// The program should allow users to input scores for each subject, calculate the average score, and assign grades based on predefined grading criteria.

// Use scanIn, Printf, loops, maps

package main

import (
	"fmt"
)

func main() {
	// Define the subjects and their weights in the final grade calculation
	subjects := []string{"Math", "English", "Science"}
	weights := map[string]float64{
		"Math":    0.4,
		"English": 0.3,
		"Science": 0.3,
	}

	// Create a map to store the scores for each subject
	scores := make(map[string]float64)

	// Prompt user to input scores for each subject
	for _, subject := range subjects {
		fmt.Printf("Enter the score for %s: ", subject)
		var score float64
		fmt.Scanln(&score)
		scores[subject] = score
	}

	// Calculate the average score based on weights
	var totalWeightedScore float64
	for subject, score := range scores {
		totalWeightedScore += score * weights[subject]
	}

	// Print the final grade
	fmt.Printf("Average Score: %.2f\n", totalWeightedScore)

	// Assign grades based on predefined grading criteria
	grade := calculateGrade(totalWeightedScore)
	fmt.Printf("Final Grade: %s\n", grade)
}

// Function to calculate the final grade based on the average score
func calculateGrade(averageScore float64) string {
	switch {
	case averageScore >= 90:
		return "A+"
	case averageScore >= 80:
		return "A"
	case averageScore >= 70:
		return "B+"
	case averageScore >= 60:
		return "B"
	case averageScore >= 50:
		return "C"
	case averageScore >= 40:
		return "D"
	case averageScore >= 30:
		return "F"
	case averageScore >= 20:
		return "FAIL"
	default:
		return "FAIL"
	}
}
