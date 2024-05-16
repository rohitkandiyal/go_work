package main

import (
	"errors" // can be used to set what err mssg to pass
	"fmt"
	"strconv"
)

func main() {
	//Converting ascii to int with error handling
	i, err := strconv.Atoi("42b")
	if err != nil { // catching error
		fmt.Println("Could not convert", err)
		return // exit main here only
	}
	fmt.Println(i) // will reach here only if successful
}

// Std lib has an interface named error with method signature Error()--  So we can create our own custom type error by jst implementing this struct.
