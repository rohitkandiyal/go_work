package main

import (
	"fmt"
)

func main() {

	//	//////////////////////// IF - ELSE //////////////////
	num1 := 20
	num2 := 20

	if num1 < num2 {
		fmt.Printf("smaller\n")
	} else if num1 == num2 {
		fmt.Printf("equal\n")
	} else {
		fmt.Printf("bigger\n")
	}

	if num3 := 35; num3 < 40 { // SCOPE of num3 is within this if block.
		fmt.Printf("kasmdkladm\n")
	}

	//	//////////////////////// FOR //////////////////
}
