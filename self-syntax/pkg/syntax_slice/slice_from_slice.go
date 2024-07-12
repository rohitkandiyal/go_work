package main

import (
	"fmt"
)

func main() {

	// Disadv of not using make
	names := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	allNames := names[1:]
	someNames := names[1:3]

	fmt.Println("len allNames:", len(allNames))
	fmt.Println("cap allNames:", cap(allNames))

	fmt.Println("allNames:", allNames)
	fmt.Println("someNames:", someNames)

	allNames = append(allNames, "Gloves")
	allNames[1] = "Canoe"

	fmt.Println("allNames:", allNames)

}
