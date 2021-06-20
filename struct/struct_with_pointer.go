package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name:", (*emp8).firstName)
	//fmt.Println("First Name Address:", emp8)
	fmt.Println("Age:", (*emp8).age)
}
