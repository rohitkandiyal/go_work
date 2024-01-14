package main

import (
	"fmt";
	"slices";
)

func main()  {
	var a []int
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a == nil) 		// empty check

	// Compare
	fmt.Println("######Compare######")
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false

	// len and cap and append
	fmt.Println("######Len-Cap######")
	var b []int
	fmt.Println(b, len(b), cap(b))
	b = append(b, 10)
	fmt.Println(b, len(b), cap(b))

	// emptying slice
	fmt.Println("######Len-Cap######")
	fmt.Println(b, len(b))
	clear(b)
	fmt.Println(b, len(b))
}
