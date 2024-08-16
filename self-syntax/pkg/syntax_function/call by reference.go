package main

import (
	"fmt"
)

// Call by value -- This is DEFAULT BEHAVIOUR
func swapNumbers(first, second int) {
	fmt.Println("Before swap:", first, second)

	var temp int
	temp = first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

// Call by reference
func swapNumbersByRef(first, second *int) {
	fmt.Println("Before swap:", *first, *second)

	var temp int
	temp = *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func main() {
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapNumbers(val1, val2)
	fmt.Println("After calling function", val1, val2)

	fmt.Println("_________ CALL BY REFERENCE _________")

	val3, val4 := 50, 60
	fmt.Println("Before calling function", val3, val4)
	swapNumbersByRef(&val3, &val4)
	fmt.Println("After calling function", val3, val4)

}
