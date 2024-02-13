// This Program was created by Aravind Nandakumar(2347214) III MCA B

// You're tasked with building a student information system in GoLang for a
// school. Each student record needs to store details such as student ID, name, age,
// and grade. Define variables to store the information of a single student and
// assign values accordingly. Pay attention to selecting appropriate data types to
// represent each piece of information.
package main

import "fmt"

// Function to add student details
func addDetails() Student {
	var student Student

	fmt.Println("Enter Student Details:")
	fmt.Print("ID: ")
	fmt.Scan(&student.ID)
	fmt.Print("Name: ")
	fmt.Scan(&student.Name)
	fmt.Print("Age: ")
	fmt.Scan(&student.Age)
	fmt.Print("Grade: ")
	fmt.Scan(&student.Grade)

	return student
}

// Function to display student details
func displayDetails(student Student) {
	fmt.Printf("[ID: %s]  ", student.ID)
	fmt.Printf("[Name: %s]  ", student.Name)
	fmt.Printf("[Age: %d]  ", student.Age)
	fmt.Printf("[Grade: %s]  ", student.Grade)
	fmt.Println()
}

type Student struct {
	ID    string
	Name  string
	Age   uint8
	Grade string
}

func main() {
	var choice int32
	choice = 10
	// a struct variable in go to store student details
	var students []Student

	for choice != 0 {
		fmt.Println("\n-----Student Information System-----")
		fmt.Println("Please select an option: ")
		fmt.Println("1. Add Student Details")
		fmt.Println("2. Display Student Details")
		fmt.Println("0. Quit")
		fmt.Scan(&choice)
		switch choice {
		// Add student details and append to the slice
		case 1:
			student := addDetails()
			students = append(students, student)

		//Display Students
		case 2:
			fmt.Print("-----All Students-----\n")
			for _, student := range students {
				displayDetails(student)
			}
			fmt.Print("----------------------\n")
		case 0:
			break
		}
	}

}
