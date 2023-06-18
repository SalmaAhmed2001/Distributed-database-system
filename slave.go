package main

import (
	"fmt"
	"io/ioutil"
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

		switch string(request[:n]) {
		case "get_student_data":
			// Read student data from file
			data, err := ioutil.ReadFile("student_data.txt")
			if err != nil {
				fmt.Println("Error reading student data from file:", err)
				return
			}

			// Send response to client
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println("Error sending response to client:", err)
				return
			}
		case "get_teacher_data":
			// Read teacher data from file
			data, err := ioutil.ReadFile("teacher_data.txt")
			if err != nil {
				fmt.Println("Error reading teacher data from file:", err)
				return
			}

			// Send response to client
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println("Error sending response to client:", err)
				return
			}
		default:
			// Invalid request
			_, err = conn.Write([]byte("Invalid request"))
			if err != nil {
				fmt.Println("Error sending response to client:", err)
				return
			}
		}
	}
}

func main() {
	// Start slave server
	listener, err := net.Listen("tcp", "localhost:58787")
	if err != nil {
		fmt.Println("Error starting slave server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Slave server started on localhost:58787")

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
