// TASK 2
// Generate the Fibonacci series up to a given number recursively.
//  The Fibonacci series is a sequence of numbers where each number
//  is the sum of the two preceding ones, starting from 0 and 1.

package main

import "fmt"

// Recursive function to generate Fibonacci series
func fibonacci(n int) []int {
	series := make([]int, 0)
	// Generate Fibonacci series up to n
	for i := 0; ; i++ {
		if fib := fibonaciRecursive(i); fib <= n {
			series = append(series, fib)
		} else {
			break
		}
	}
	return series
}

// Recursive helper function to calculate Fibonacci number at index n
func fibonaciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaciRecursive(n-1) + fibonaciRecursive(n-2)
}

func main() {
	// Input from the user
	var limit int
	fmt.Print("Enter the limit for Fibonacci: ")
	fmt.Scan(&limit)

	// Generate Fibonacci series up to the given limit
	series := fibonacci(limit)

	// Print the series
	fmt.Println("Fibonacci series up to", limit, ":", series)
}
