package main

import (
	"fmt"

	"github.com/rohitkandiyal/go_work/pkg/test_package"

	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("main.go: Hello from go_work module main...!")
			test_package.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()

}
