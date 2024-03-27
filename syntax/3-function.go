package main

import (
	"fmt"
)

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

func add(x , y int) int {
	return x + y
}

// Ideally in go all the params to a func should be provided. But if needed then below are the ways for named and optional params

// Function with NAMED PARAMS - done using a struct.
// Function with optional params - done using ...


func main()  {
	fmt.Println(concat("hellooo", "rohit..."))
	fmt.Println(add(2,3))
}