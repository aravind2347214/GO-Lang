package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 100, "medium": 200, "large": 400}},
		{name: "Espresso", prices: map[string]float64{"single": 100, "double": 200}},
	}

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
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Enter a new menu item")
			item, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: item, prices: make(map[string]float64, 0)})

		case "q":
			break loop

		default:
			fmt.Println("Please enter the correct option")
		}
	}
}
