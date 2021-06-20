package main

import (
	"fmt"
)

func main() {

	a := hello()
	fmt.Println("a of address ", a)
	fmt.Println("a of val ", *a)
}

func hello() *int {
	a := 10
	return &a
}
