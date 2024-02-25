package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var menu = []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 100, "medium": 200, "large": 400}},
	{name: "Espresso", prices: map[string]float64{"single": 100, "double": 200}},
}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}

}

func itemExists(itemName string) bool {
	for _, item := range menu {
		if strings.EqualFold(item.name, itemName) {
			return true
		}
	}
	return false
}

func addItem() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the name of the new item:")
	item, _ := in.ReadString('\n')
	item = strings.TrimSpace(item)

	if itemExists(item) {
		fmt.Printf("Item '%s' already exists in the menu.\n", item)
	} else {
		menu = append(menu, menuItem{name: item, prices: make(map[string]float64)})
		fmt.Printf("Item '%s' added to the menu.\n", item)
	}
}

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
			printMenu()
		case "2":
			addItem()
		case "q":
			break loop

		default:
			fmt.Println("Please enter the correct option")
		}
	}
}
