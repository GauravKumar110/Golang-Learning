package main

import (
	"fmt"
)

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", b)
	fmt.Println("address of b is", a)
	fmt.Println("address of b is", b)
	fmt.Println("address of b is", &b)
}
