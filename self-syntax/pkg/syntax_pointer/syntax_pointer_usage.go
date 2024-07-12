package main

import (
	"fmt"
	"sort"
)

func main() {

	// Not using pointer - Pass By Value
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]

	fmt.Println(secondName)

	sort.Strings(names[:])

	fmt.Println(names)
	fmt.Println(secondName)

	// Not using pointer - Pass By Reference
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")

	newNames := [3]string{"Alice", "Charlie", "Bob"}

	var secondPosition *string
	secondPosition = &newNames[1]

	fmt.Println(*secondPosition)
	sort.Strings(newNames[:])

	fmt.Println(newNames)
	fmt.Println(*secondPosition)
}
