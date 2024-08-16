package main

import (
	"fmt"
)

// Defining alias to func signature
type calcFunc func(float64) float64

// printPrice accepts 3 args - string/float64/function
func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// selectCalculator takes floats and return function

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		// here we are using literal or anonymous func
		var withTax calcFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return withTax
	}
	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}
