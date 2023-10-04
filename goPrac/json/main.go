package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string
	Email    string
	password string
	courses  []string
}

func main() {
	fmt.Println("here we will see encode/decode of json")
	encodeJson()

}

func encodeJson() {
	users := []user{
		{"Naveen", "n@g.com", "1234", []string{"java", "js"}},
		{"Kaushal", "k@g.com", "9876", []string{"react", "js"}},
	}

	//package this data as json data

	finalJson, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}
