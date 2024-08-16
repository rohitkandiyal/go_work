package main

import (
	"fmt"
)

func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers { // THis means values are always received as slice and therefore we can enumerate them.
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func main() {
	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Lifejacket", "Sail Safe Co")

	printSuppliers("Soccer Ball") // empty variadic param

	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Kayak", names...) // passing slice with values as variadic param
}
