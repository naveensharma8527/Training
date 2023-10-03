package main

import (
	"fmt"
	"sort"
)

func main() {

	var names = []string{}

	names = append(names, "nav", "raj", "shu", "kau")

	fmt.Println("list: ", names)

	// names = append(names[1:])
	sort.Strings(names)

	fmt.Println("updated list: ", names)

	// remove a value from slice based on index

	var alp = []string{"a", "b", "c", "d", "e"}

	var ind int = 1

	alp = append(alp[:ind], alp[ind+1:]...)

	fmt.Println(alp)

}
