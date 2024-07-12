package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	m1 := map[string]int{
		"hello":  5,
		"duniya": 0,
	}

	//without comma ok
	fmt.Println(m["world"])
	fmt.Println(m1["world"]) // ERROR: it will also give 0(zero value of int type) but should not ideally as m1 doesn't have world key.

	// Comma ok idiom
	fmt.Println("+++++++++++Comma ok idiom+++++++++++++")

	v, ok := m["world"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key not found")
	}

	v, ok = m1["world"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key not found")
	}

}
