package main

import (
	"errors"
	"fmt"
)

// write me a code to illustrate the example of errors in go language
func main() {
	var word string
	fmt.Println("Enter word (If error Entered or less than 5 error will trigger)")
	fmt.Scan(&word)
	err := doSomething(word)
	if err != nil {
		fmt.Println(err) // prints "error: hello is not good enough!"
	} else {
		fmt.Println("No error all Good")
	}
}

func doSomething(s string) error {
	if s == "error" || len(s) < 5 {
		return errors.New("An error has Occured")
	}
	return nil
}
