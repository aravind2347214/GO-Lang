package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {

	}()

	msg := <-messages
	fmt.Println(msg)
}
