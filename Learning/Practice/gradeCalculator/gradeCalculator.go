// You are tasked with creating a grade calculator program in GoLang. The
// program should prompt the user to input their scores in three subjects: Math,
// Science, and English. Based on these scores, calculate the average grade
// (assuming each subject is equally weighted) and display the corresponding letter
// grade (A, B, C, D, or F) using control flow.

package main

import (
	"fmt"
)

// Subject struct represents a subject with its name and score.
type Subject struct {
	Name  string
	Score float64
}

func main() {
	// Create subjects for Math, Science, and English
	math := Subject{Name: "Math", Score: getUserInput("Enter Math score: ")}
	science := Subject{Name: "Science", Score: getUserInput("Enter Science score: ")}
	english := Subject{Name: "English", Score: getUserInput("Enter English score: ")}

	// Calculate average score
	averageScore := calculateAverage(math.Score, science.Score, english.Score)

	// Display average score and corresponding letter grade
	displayResult(averageScore)
}

// getUserInput takes user input for subject scores and returns a float64.
func getUserInput(prompt string) float64 {
	var score float64
	fmt.Print(prompt)
	fmt.Scan(&score)
	return score
}

// calculateAverage calculates the average score for three subjects.
func calculateAverage(math, science, english float64) float64 {
	return (math + science + english) / 3.0
}

// displayResult displays the average score and corresponding letter grade.
func displayResult(averageScore float64) {
	fmt.Printf("\nAverage Score: %.2f\n", averageScore)

	// Determine the letter grade based on the average score
	letterGrade := getLetterGrade(averageScore)
	fmt.Printf("Letter Grade: %s\n", letterGrade)
}

// getLetterGrade determines the letter grade based on the average score.
func getLetterGrade(averageScore float64) string {
	switch {
	case averageScore >= 90.0:
		return "A"
	case averageScore >= 80.0:
		return "B"
	case averageScore >= 70.0:
		return "C"
	case averageScore >= 60.0:
		return "D"
	default:
		return "F"
	}
}
