package main

import (
	"fmt"
)

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		string: "naveen",
		int:    50,
	}
	p2 := Person{"GAURAV", 35}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	fmt.Println(p2)
}
