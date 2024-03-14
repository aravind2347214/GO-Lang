package main

import "fmt"

func main() {

	a := 10
	b := 20
	fmt.Printf("A = %d and B = %d", a, b)
	swap(&a, &b)
	fmt.Println("\nAfter Swap")
	fmt.Printf("A = %d and B = %d", a, b)

}

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
