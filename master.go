package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		// Handle client request
		request := make([]byte, 1024)
		n, err := conn.Read(request)
		if err != nil {
			fmt.Println("Error reading request from client:", err)
			return
		}

		// Forward request to slave server
		slaveConn, err := net.Dial("tcp", "localhost:58787")
		if err != nil {
			fmt.Println("Error connecting to slave:", err)
			return
		}
		defer slaveConn.Close()

		_, err = slaveConn.Write(request[:n])
		if err != nil {
			fmt.Println("Error forwarding request to slave:", err)
			return
		}

		// Read response from slave server
		response := make([]byte, 1024)
		n, err = slaveConn.Read(response)
		if err != nil {
			fmt.Println("Error reading response from slave:", err)
			return
		}

		// Send response to client
		_, err = conn.Write(response[:n])
		if err != nil {
			fmt.Println("Error sending response to client:", err)
			return
		}
	}
}

func main() {
	// Start master server
	listener, err := net.Listen("tcp", "localhost:56588")
	if err != nil {
		fmt.Println("Error starting master server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Master server started on localhost:56588")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle client request in separate goroutine
		go handleClient(conn)
	}
}
