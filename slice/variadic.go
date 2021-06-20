package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("a: ", a, "has hold length", len(a), "and capacity", cap(a))
	a = append(a, 8, 9, 10, 11, 12, 13)
	fmt.Println("a: ", a, "has hold length", len(a), "and capacity", cap(a))
}
