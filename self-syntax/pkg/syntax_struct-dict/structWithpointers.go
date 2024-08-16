// Structs here are like dicts. But we can create methods(check 4 below) on a struct and in that way it behaves like a class.
// As it can have attributes and methods both.

package main

import (
	"fmt"
)

func main() {
	// Struct without pointer: Pass by value

	type Product struct {
		name, category string
		price          float64
	}

	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p2 := p1

	p1.name = "Original Kayak"

	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
	// Struct with pointer: Pass by ref
	p3 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	p4 := &p3

	p3.name = "Original Kayak"

	fmt.Println("P1:", p3.name)
	fmt.Println("P2:", (*p4).name)

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
	// Defining pointer with new keyword
	var p5 *Product
	p5 = new(Product)
	fmt.Println("P5:", p5)
	fmt.Println("P5:", &p5)
	fmt.Println("P5:", *p5)
}
