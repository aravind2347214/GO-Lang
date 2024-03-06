package main

import (
	"bufio"
	"fmt"
	"os"
	"sample-app-2/conditionals"
	"strings"
)

// type MenuItem struct {
// 	name   string
// 	prices map[string]float64
// }

func main() {

loop:
	for {
		fmt.Println("Please Enter your Choice")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		in := bufio.NewReader(os.Stdin)
		choice, _ := in.ReadString('\n')
		switch strings.TrimSpace(choice) {
		case "1":
			conditionals.PrintMenu()
		case "2":
			conditionals.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Please enter the correct option")
		}
	}
}
