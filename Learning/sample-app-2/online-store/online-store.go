// This Program was created by Aravind Nandakumar(2347214) III MCA B

// You're developing an online store application in GoLang. As part of the
// application, you need to keep track of various product details such as name,
// price, and quantity in stock. Design a set of variables and assign values to
// represent a specific product in the inventory. Ensure you use appropriate data
// types for each variable to accurately capture the information.
package main

import "fmt"

// Struct to define a product
type Product struct {
	productName string
	price       float64
	quantity    int32
}

// function to add products
func addProduct() Product {
	var product Product
	fmt.Println("Enter Product Details:")
	fmt.Print("Product Name: ")
	fmt.Scan(&product.productName)
	fmt.Print("Price: ")
	fmt.Scan(&product.price)
	fmt.Print("Quantity: ")
	fmt.Scan(&product.quantity)
	return product
}

// Function to display product details
func displayDetails(product Product) {
	fmt.Printf("[Product Name: %s]  ", product.productName)
	fmt.Printf("[Price: %0.3f]  ", product.price)
	fmt.Printf("[Quantity: %d]  ", product.quantity)
	fmt.Println()
}

func main() {
	// choice variable to give user choice capability
	var choice int32
	choice = 10
	// a struct variable in go to store Product details
	var products []Product

	for choice != 0 {
		fmt.Println("\n-----Online Store-----")
		fmt.Println("Please select an option: ")
		fmt.Println("1. Add Product")
		fmt.Println("2. Display All Products")
		fmt.Println("0. Quit")
		fmt.Scan(&choice)
		switch choice {
		// Add product details and append to the slice
		case 1:
			product := addProduct()
			products = append(products, product)

		//Display Products
		case 2:
			fmt.Print("-----All Products-----\n")
			for _, product := range products {
				displayDetails(product)
			}
			fmt.Print("----------------------\n")
		case 0:
			break
		}
	}

}
