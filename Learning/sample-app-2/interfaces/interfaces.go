package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

func (u user) Print() string {
	return fmt.Sprintf("The user name and id is %v  : [%v]", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

func main() {
	var p printer
	p = user{username: "Shiva", id: 6}
	fmt.Println(p.Print())
	p = menuItem{name: "Coffee", prices: map[string]float64{"small": 200, "medium": 300, "large": 600}}
	fmt.Println(p.Print())
	u, ok := p.(user)
	fmt.Println(u, ok)
	u1, ok := p.(menuItem)
	fmt.Println(u1, ok)

	switch v := p.(type) {
	case user:
		fmt.Println("User Found:", v)
	case menuItem:
		fmt.Println("Menu Item:", v)

	default:
		fmt.Println("Unknown type")

	}

}
