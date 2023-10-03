package main

import "fmt"

func main() {

	naveen := Student{1, "Full Stack"}

	fmt.Println(naveen.getinfo())

}

type Student struct {
	Roll   int
	Course string
}

func (st Student) getinfo() string {
	return "Course is " + st.Course
}
