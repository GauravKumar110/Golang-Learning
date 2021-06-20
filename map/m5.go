//delete valu from the map
package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	//iterate map
	for i, v := range employeeSalary {
		fmt.Println("Salary of", i, "is", v)
	}

	delete(employeeSalary, "mike")
	for i, v := range employeeSalary {
		fmt.Println("New salary of", i, "is", v)
	}
}
