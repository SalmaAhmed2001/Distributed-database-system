package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: client <server> <request>")
		os.Exit(1)
	}

	server := os.Args[1]
	request := os.Args[2]

	// Connect to server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send request to server
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error sending request to server:", err)
		return
	}

	// Read response from server
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error reading response from server:", err)
		return
	}

	fmt.Println(string(response[:n]))
}
