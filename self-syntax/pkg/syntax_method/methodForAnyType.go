// Structs here are like dicts. But we can create methods(check 4 below) on a struct and in that way it behaves like a class.
// As it can have attributes and methods both.

package main

import (
	"fmt"
)

type Name []string

// STRUCT Method for alias Name
func (name *Name) printName() {
	for _, p := range *name {
		fmt.Println("Name is:", p)
	}
}

func main() {
	nameSlice := Name{"Rohit", "KAnDiyal"}
	nameSlice.printName()
}
