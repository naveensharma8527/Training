package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections on port 9988
	listen, err := net.Listen("tcp", ":9988")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server is listening on port 9988")

	for {
		// Accept a client connection
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle the client's request in a new goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read client's message
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Process the message (you can add your custom logic here)
	clientMessage := string(buffer)
	fmt.Println("Received message from client:", clientMessage)

	// Send a response back to the client
	response := "Hello from the server!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Println("Sent response to client:", response)
}
