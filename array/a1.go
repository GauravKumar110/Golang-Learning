package main

import (
	"fmt"
)

func main() {
	var a [3]bool //int array with length 3
	fmt.Println(a)

	b := [...]int8{1, 2, 3, 4, 5, 6} //int array with length 6
	fmt.Println(len(b))
	fmt.Println(b)
}
