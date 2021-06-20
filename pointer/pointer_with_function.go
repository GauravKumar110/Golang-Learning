package main

import (
	"fmt"
)

func main() {

	a := 10

	fmt.Println("Starting value = ", a)
	b := &a
	change(b)
	fmt.Println("after update value = ", a)
}

func change(val *int) {
	*val = 15
}
