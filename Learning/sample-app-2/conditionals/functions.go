package conditionals

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Menu represents a restaurant menu.
type Menu struct {
	items []MenuItem
}

// NewMenu creates a new menu.
func NewMenu() *Menu {
	return &Menu{
		items: []MenuItem{
			{name: "Coffee", prices: map[string]float64{"small": 100, "medium": 200, "large": 400}},
			{name: "Espresso", prices: map[string]float64{"single": 100, "double": 200}},
		},
	}
}

// MenuItem represents a menu item.
type MenuItem struct {
	name   string
	prices map[string]float64
}

// itemExists checks if an item already exists in the menu.
func (m *Menu) itemExists(itemName string) bool {
	for _, item := range m.items {
		if strings.EqualFold(item.name, itemName) {
			return true
		}
	}
	return false
}

// AddItem adds a new item to the menu.
func (m *Menu) AddItem() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the name of the new item:")
	item, _ := in.ReadString('\n')
	item = strings.TrimSpace(item)

	if m.itemExists(item) {
		fmt.Printf("Item '%s' already exists in the menu.\n", item)
	} else {
		m.items = append(m.items, MenuItem{name: item, prices: make(map[string]float64)})
		fmt.Printf("Item '%s' added to the menu.\n", item)
	}
}

// PrintMenu prints the entire menu.
func (m *Menu) PrintMenu() {
	for _, item := range m.items {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
