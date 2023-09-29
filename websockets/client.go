package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server at localhost:9988
	conn, err := net.Dial("tcp", "localhost:9988")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Read a custom message from the user
	fmt.Print("Enter a message to send to the server: ")
	reader := bufio.NewReader(os.Stdin)
	customMessage, _ := reader.ReadString('\n')

	// Send the custom message to the server
	_, err = conn.Write([]byte(customMessage))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	fmt.Println("Sent custom message to server:", customMessage)

	// Receive and print the response from the server
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	serverResponse := string(buffer)
	fmt.Println("Received response from server:", serverResponse)
}
