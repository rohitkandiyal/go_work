// Strings are immutable like Python

package main

import (
	"fmt"
)

func main() {
	str1 := "Greetings and\n\"Salutations\""

	// raw strings is between ``
	str2 := `Greetings and
	"Salutations"`

	fmt.Printf("This is str1: %s", str1)
	fmt.Printf("\n###########\n")
	fmt.Printf("This is str2: %s\n", str2)

	// Concatenation
	f_name := "Rohit"
	l_name := "Kandiyal"
	fmt.Printf("Concatenation: %s\n", f_name+" "+l_name)

	// str working
	fmt.Printf("\n####int str working####\n")
	s := "Hello there"
	b := s[6]
	fmt.Println(b) // prints 116 which is UTF-8 encoding of "t"

}
