package main

import (
	"fmt"

	"github.com/rohitkandiyal/go_work/pkg/test_package"
)

func main() {

	fmt.Println("main.go: Hello from go_work module main...!")
	test_package.PrintHello()

}
