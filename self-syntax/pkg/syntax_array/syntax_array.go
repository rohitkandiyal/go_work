package main

import "fmt"

func main() {

	var names [3]string
	fmt.Println(names) // Check zero value of ARRAY
	// fmt.Println(names == nil) // empty check - check why it fails

	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"

	fmt.Println(names)

	// Changing default behaviour of pass by value in arrays to pass by ref using pointers
	fmt.Println("++++++++++++++++++++++++++++++++++")

	newNames := [3]string{"Kayak", "Lifejacket", "Paddle"}

	otherNames := &newNames

	newNames[0] = "rohit"
	fmt.Println("newName:", newNames)
	fmt.Println("otherNames:", *otherNames)

	fmt.Println("++++++++++++++++++++++++++++++++++")
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}
}
