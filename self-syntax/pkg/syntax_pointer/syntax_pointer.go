package main

import (
	"fmt"
)

func main() {

	first := 100
	second := &first

	first++
	*second++

	var myNewPointer *int
	fmt.Println("myNewPointer Zero value for mem location:", myNewPointer)
	// fmt.Println("myNewPointer Zero value for deref value:", *myNewPointer)		// This fails as we can't deref a pointer with nil value
	myNewPointer = second
	*myNewPointer++

	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Second:", *second)

	// Pointing at pointer example:

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")
	third := 20
	fourth := &third
	fifth := &fourth // fifth is pointer to pointer fourth and is allowed in GO

	fmt.Println("Fourth:", fourth)
	fmt.Println("Value at pointer Fourth:", *fourth)
	fmt.Println("Fifth:", fifth)
	fmt.Println("Value at pointer Fifth:", *fifth)
	fmt.Println("value of Value at pointer Fifth:", **fifth)
}
