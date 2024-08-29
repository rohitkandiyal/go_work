package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

// type method
func (p *Product) printProductDetails() {
	fmt.Println("Name:", p.name, "Category:", p.category,
		"Price", p.price)
}

// Methods implementing interface
func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}
