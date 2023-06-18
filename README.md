# Distributed-database-system
TCP-based Master-Slave Client-Server Application for Sharing Student and Teacher Data using the Go language 

The first program creates a TCP client that sends a request to a server and prints the response received from the server.

The second program creates a TCP proxy server that forwards client requests to a slave server and sends back the response to the client. 

The third program creates a TCP file server that serves student and teacher data in response to specific client requests.

Together, these programs demonstrate how to build a simple distributed system using Go's net package and how to implement a master-slave architecture.
