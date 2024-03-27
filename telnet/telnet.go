package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run telnet_client.go <host> <port>")
		os.Exit(1)
	}

	host := os.Args[1]
	port := os.Args[2]

	// Connect to the server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", host+":"+port)

	// Start a goroutine to read and print responses from the server
	go readResponses(conn)

	// Get user input and send it to the server
	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	fmt.Print("Enter command (type 'exit' to quit): ")
	// 	scanner.Scan()
	// 	command := scanner.Text()

	// 	if command == "exit" {
	// 		break
	// 	}

	// 	// Send the command to the server
	// 	fmt.Fprintln(conn, command)
	// }
}

func readResponses(conn net.Conn) {
	// Create a reader to read responses from the server
	reader := bufio.NewReader(conn)

	for {
		// Read the response from the server
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			os.Exit(1)
		}

		// Print the response
		fmt.Print("Server response: " + response)
	}
}
