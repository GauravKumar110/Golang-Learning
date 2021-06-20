//Retrieving value for a key from a map
//iterate map
//check key found or not in map
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
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	//iterate map
	for i, v := range employeeSalary {
		fmt.Println("Salary of", i, "is", v)
	}

	emp := "mike"
	fmt.Println("Salary of", emp, "is", employeeSalary[emp])
	value, ok := employeeSalary[emp]
	fmt.Println(value, ok)
	//check key found or not in map
	if ok {
		fmt.Println("Salary of", emp, "is", value)
	}
}
