// TASK 1
// You are given a list of integers. Your task is to find all the pairs of integers in
// the list whose sum is equal to a given target value. However, each integer in
// the list can only be used once in a pair. Write a Go function find Pairs that
// takes in a list of integers numbers, and an integer target, and returns a slice
// of pairs of integers that sum up to the target. Array elements and the target
// should be given by the user

// Suppose the array elements are 2, 7, 11, 15, 3, 6, 8, 12    and the target value is  18, the expected output is

// Pairs with sum 18 are: [[7 11] [6 12]]

package main

import (
	"fmt"
	"sort"
)

func calculatePairs(numbers []int, target int) [][]int {
	sort.Ints(numbers)
	var pairs [][]int

	// Use two pointers technique
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			pairs = append(pairs, []int{numbers[left], numbers[right]})
			// Move both pointers to find other pairs
			left++
			right--
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return pairs
}

func main() {
	// Input from the user
	numbers := []int{5, 5, 7, 3, 6, 2, 8, 5, 1, 9}
	target := 10

	// Find pairs
	pairs := calculatePairs(numbers, target)

	// Print result
	fmt.Printf("Pairs with sum %d are: %v\n", target, pairs)
}
