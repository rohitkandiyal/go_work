package test_package

import "fmt"

func PrintHello() {
	fullName := getName()
	fmt.Printf("Hello, %s ! This is test_package speaking! \n", fullName)
}
