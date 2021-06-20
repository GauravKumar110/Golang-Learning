package main

import (
	"fmt"
)

func main() {
	b := 255
	a := &b
	c := &a
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	fmt.Println("value of b is", *c)
	*a++
	fmt.Println("new value of b is", b)
	fmt.Println("new value of b is", a)
}
