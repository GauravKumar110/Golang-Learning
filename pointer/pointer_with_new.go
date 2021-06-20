package main

import (
	"fmt"
)

func main() {
	size := new(int)
	c := 100

	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	size = &c
	fmt.Println("New size value is", *size)

	b := 255
	a := new(int)
	a = &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}
