package main

import "fmt"

type emp struct {
	name, city string
	age        int8
}

func main() {
	emp1 := emp{
		"gaurav", "haridwar", 35,
	}

	fmt.Println(emp1)
}
