package main

import "fmt"

type User struct {
	Name string
	age  int
}

func main() {

	fmt.Println("Structs ...")

	naveen := User{"naveen", 24}

	fmt.Println(naveen)

	//print complete structs
	fmt.Printf("Naveen details : %+v\n", naveen)

}
