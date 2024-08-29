package main

import (
	"fmt"
)

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}

	// Without interface ---  and THIS WILL ALSO WORK
	fmt.Println("+++++++++Without interface+++++++++")
	kayak.printProductDetails()
	insurance.printServiceDetails()
	fmt.Println("Total expense incurred:", kayak.price+(insurance.monthlyFee*float64(insurance.durationMonths)))

	// With interface
	fmt.Println("+++++++++With interface+++++++++")
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

}
