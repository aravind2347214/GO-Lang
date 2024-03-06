// TASK 4
// Generate prime numbers up to a given limit  Prime numbers are numbers
//
//	greater than 1 that have no positive divisors other than 1 and themselves
package main

import (
	"fmt"
)

func getPrimes(limit int) []int {
	if limit < 2 {
		return nil
	}

	// Initialize a slice to store prime numbers
	primes := []int{}

	// Create a boolean slice to mark non-prime numbers
	sieve := make([]bool, limit+1)

	// Start with 2, the first prime number
	for i := 2; i <= limit; i++ {
		// If the number is not marked as non-prime, it's a prime number
		if !sieve[i] {
			primes = append(primes, i)
			// Mark all multiples of the prime number as non-prime
			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}

	return primes
}

func main() {
	// Input from the user
	var limit int
	fmt.Print("Enter the limit for primes: ")
	fmt.Scan(&limit)

	// Generate prime numbers up to the given limit
	primes := getPrimes(limit)

	// Print the prime numbers
	fmt.Printf("Prime numbers up to %d are : %v\n", limit, primes)
}
