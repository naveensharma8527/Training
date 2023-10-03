package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome user "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name")

	// comma ok || err ok syntax

	input, _ := reader.ReadString('1')

	fmt.Println("welcome ", input)

}
