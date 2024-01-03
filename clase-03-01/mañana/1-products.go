package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Apple",
		Price:       1.99,
		Description: "Red apple",
		Category:    "Fruit",
	},
	{
		ID:          2,
		Name:        "Orange",
		Price:       1.75,
		Description: "An orange fruit",
		Category:    "Fruit",
	},
	{
		ID:          3,
		Name:        "Shirt",
		Price:       20.99,
		Description: "A white shirt",
		Category:    "Clothing",
	},
}

func main() {
	pants := Product{
		ID:          4,
		Name:        "Pants",
		Price:       25.99,
		Description: "A pair of pants",
		Category:    "Clothing",
	}

	pants.Save()

	pants.GetAll()

	product := GetById(1)
	fmt.Println(product)
}

func (product Product) Save() {
	Products = append(Products, product)
	return
}

func (product Product) GetAll() {
	for _, p := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %f, Description: %s, Category: %s\n", p.ID, p.Name, p.Price, p.Description, p.Category)
	}
	return
}

func GetById(id int) Product {
	for _, p := range Products {
		if p.ID == id {
			return p
		}
	}
	return Product{}
}
