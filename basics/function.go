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

func main()  {
	fmt.Println(concat("hellooo", "rohit..."))
	fmt.Println(add(2,3))
}