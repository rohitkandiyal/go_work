package main

import (
	"fmt"
)

func main() {

	// Disadv of not using make
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	pointerA := &names[0]

	fmt.Println(names)
	fmt.Println(pointerA)

	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))

	// Append to slice nt created via make

	names = append(names, "Hat", "Gloves")
	pointerB := &names[0]

	fmt.Println(names)
	fmt.Println(pointerB)

	// Adv of using make
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")

	newNames := make([]string, 3, 6)
	pointerC := &newNames[0]

	newNames[0] = "Kayak"
	newNames[1] = "Lifejacket"
	newNames[2] = "Paddle"

	fmt.Println(newNames)
	fmt.Println(pointerC)

	fmt.Println("len:", len(newNames))
	fmt.Println("cap:", cap(newNames))

	newNames = append(newNames, "Hat", "Gloves")
	pointerD := &newNames[0]

	fmt.Println(newNames)
	fmt.Println(pointerD)

}
