package main

import (
	"fmt"
)

func main()  {

	var firstname string = "rohit"
	lastname := "kandiyal"				// Dynamic assignment
	age := 35

	fmt.Printf("My name is %s %s \nMy age is %d\n", firstname, lastname, age)
	fmt.Printf("Age is of type %T\n", age)
	
	// Type casting in go needs explicit conversion as go is strongly typed

	// a := "18"
	
	// fmt.Printf("%d\n", int(a)+10) 				//a can't be coverted to int from str 

	const lname = "KANDIYAL"			// No use of := in constants
	// lname = "asnkl"					// This will break compilation

	// %v can be used when we are not sure of data type..
	fmt.Printf("My name is %v %v \nMy age is %v\n", firstname, lname, age)

	fmt.Printf("My weight is %.2f kg\n", 73.4567)

	// length

	fmt.Printf("%d\n", len("bigger"))
	
}


