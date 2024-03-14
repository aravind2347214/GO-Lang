package main

import (
	"bufio"
	"fmt"
	"os"
	"sample-app-2/conditionals"
	"strings"
)

func main() {
	menu := conditionals.NewMenu()

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
			menu.PrintMenu()
		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Please enter the correct option")
		}
	}
}
