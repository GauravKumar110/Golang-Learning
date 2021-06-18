package main

import "fmt"

func main() {
	a := [5]int{2, 3, 4, 5, 6}

	fmt.Println(a)

	var b []int = a[2:4] //a[start:end] //starting from indexes 2 through value count 4 from the starting
	fmt.Println(b)
	fmt.Println("length of b : ", len(b))
}
