package main

import(
	"fmt"
)

func main()  {
	words := []string{"a", "cow", "smile", "gopher",
    "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {

		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")

		case 5:
			// wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", len(word))

		case 6, 7, 8, 9:		// Do nothing

		default:
			fmt.Println(word, "is a long word!")
		}
	}

	fmt.Println("#####################################")
	// Above the default case feature of equality is used... But we can use any other comparison in each case too
	// This is called EMPTY SWITCH--  WHY -- 0+

	words1 := []string{"hi", "salutations", "hello"}
	for _, word := range words1 {
		switch wordLen := len(word); {						// +0 -- why because there is nothing defined between ; {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}