package main

import "fmt"

func main() {
	var a, b, c bool = false, true, true
	fmt.Println(a && b, b && c)
}
