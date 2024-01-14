package main

import(
	"fmt"
)

func main()  {
	// FOR Loop in four ways are done in Go
	// Way 1
	for i:=0; i < 10 ;i++ {
		fmt.Println(i)
	}	
	
	// Now for above has three parts(initialization; condition; increment) and all these parts are optional

	// Way 2 - WHILE LOOP -- it is basically a for loop with just condition in Go - leaving initialization and increment
	fmt.Println("#####################################")
	a := 5
	b := 10
	for a <= b{
		fmt.Println(a)
		a = a + 2
	}

	// Way 3 - Infinite LOOP -- it is basically a for loop with nothing - leaving initialization;condtion and increment

	// for {
    //     fmt.Println("Hello")
    // }

	// Way 4 - for-range loop - to iterate over built-in types like slice, strings, maps
	fmt.Println("#####################################")
	fmt.Println("iterate over slice")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {			// if index is not reqd
		fmt.Println(v)
	}

	for k := range evenVals {				// if value is not needed
		fmt.Println(k)
	}

	fmt.Println("#####################################")
	fmt.Println("iterate over string")

	str := "rohit"
	for i, r := range str {
        fmt.Println(i, r, string(r))
    }


 }


