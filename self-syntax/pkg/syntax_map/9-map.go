package main

import (
	"fmt"
	"maps"
)

func main() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	fmt.Println(m, len(m))

	// Delete from map
	delete(m, "hello")
	fmt.Println(m)

	// Completely clear
	clear(m)
	fmt.Println(m, len(m))

	// Comparing maps

	p := map[string]int{
		"hello": 5,
		"world": 10,
	}
	q := map[string]int{
		"world": 10,
		"hello": 5,
	}
	fmt.Println(maps.Equal(p, q))

}
