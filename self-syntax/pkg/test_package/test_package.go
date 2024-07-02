package test_package

import "fmt"

func PrintHello() {
	fullName := getName()
	fmt.Printf("test_package: Hello, %s ! This is test_package speaking! \n", fullName)
}
