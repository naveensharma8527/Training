package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter you txt")

	input, _ := reader.ReadString('\n')

	file, err := os.Create("./hello.txt")

	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(file, input)

	if err != nil {
		panic(err)
	}

	file.Close()

}
