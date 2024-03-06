package conditionals

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

func itemExists(itemName string) bool {
	for _, item := range menu {
		if strings.EqualFold(item.name, itemName) {
			return true
		}
	}
	return false
}

func AddItem() {
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

func PrintMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}

}
