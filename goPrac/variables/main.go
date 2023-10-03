package main

import "fmt"

func main() {
	// fmt.Println("Hi Variables")
	var username string = "naveen sharma"
	fmt.Println(username)

	fmt.Printf("variable is of type: %T \n", username)

	var flag bool = true
	fmt.Println(flag)

	fmt.Printf("variable is of type: %T \n", flag)

	var smallVal uint8 = 255

	fmt.Println(smallVal)

	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloatVal float32 = 255.55121324

	fmt.Println(smallFloatVal)

	fmt.Printf("variable is of type: %T \n", smallFloatVal)

	//implicit type
	var sayHi = "Hii Bro"
	fmt.Println(sayHi)

	//no var style
	numOfUser := 3000
	fmt.Println(numOfUser)

}
