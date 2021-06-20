//initialize a map during the declaration
package main

import (
	"fmt"
)

func main() {
	empSalary := map[string]int{
		"gaurav": 500,
		"danny":  1000,
		"atul":   5000,
	}
	fmt.Println(empSalary)
}
