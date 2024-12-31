package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product

func addProduct(id int, name string, price float64, stock int) error {
	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID must be unique")
		}
	}

	inventory = append(inventory, Product{ID: id, Name: name, Price: price, Stock: stock})
	return nil
}

func updateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

func searchProductByID(id int) (*Product, error) {
	for _, product := range inventory {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func searchProductByName(name string) (*Product, error) {
	for _, product := range inventory {
		if strings.EqualFold(product.Name, name) {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

func displayInventory() {
	fmt.Println("\nInventory:")
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func sortInventoryByPrice() {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i].Price < inventory[j].Price
	})
}

func sortInventoryByStock() {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i].Stock < inventory[j].Stock
	})
}

func main() {
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product by ID")
		fmt.Println("4. Search Product by Name")
		fmt.Println("5. Display Inventory")
		fmt.Println("6. Sort Inventory by Price")
		fmt.Println("7. Sort Inventory by Stock")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter product ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter product name: ")
			var name string
			fmt.Scan(&name)
			fmt.Print("Enter product price: ")
			var price float64
			fmt.Scan(&price)
			fmt.Print("Enter product stock: ")
			var stock int
			fmt.Scan(&stock)
			if err := addProduct(id, name, price, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully.")
			}

		case 2:
			fmt.Print("Enter product ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter new stock: ")
			var stock int
			fmt.Scan(&stock)
			if err := updateStock(id, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully.")
			}

		case 3:
			fmt.Print("Enter product ID: ")
			var id int
			fmt.Scan(&id)
			if product, err := searchProductByID(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product found: %+v\n", *product)
			}

		case 4:
			fmt.Print("Enter product name: ")
			var name string
			fmt.Scan(&name)
			if product, err := searchProductByName(name); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product found: %+v\n", *product)
			}

		case 5:
			displayInventory()

		case 6:
			sortInventoryByPrice()
			fmt.Println("Inventory sorted by price.")

		case 7:
			sortInventoryByStock()
			fmt.Println("Inventory sorted by stock.")

		case 8:
			fmt.Println("Exiting the system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
