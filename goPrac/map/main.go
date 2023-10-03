package main

import "fmt"

func main() {

	ch := make(map[string]int)

	ch["a"] = 1
	ch["b"] = 2
	ch["c"] = 3
	ch["d"] = 4
	ch["e"] = 5

	fmt.Println("map : ", ch)
	fmt.Println("value of b is ", ch["b"])
	delete(ch, "c")
	fmt.Println("map : ", ch)

	//iterate map

	for key, value := range ch {
		fmt.Printf("key is %v and value is %v\n ", key, value)
	}
}
