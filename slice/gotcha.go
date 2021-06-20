package main

import (
	"fmt"
)

func change(s [2]string) {
	s[0] = "Go"
	fmt.Println("inside change method := ", s)
}

func sliceChange(s []string) {
	s[0] = "Go"
	fmt.Println("inside slice change method := ", s)
}

func main() {
	welcome := [2]string{"hello", "world"}
	change(welcome)
	fmt.Println("new array := ", welcome)
	//passing slice
	slice := []string{"hello", "world"}
	sliceChange(slice)
	fmt.Println("new slice : = ", slice)
}
