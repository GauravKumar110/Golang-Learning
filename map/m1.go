//for store key and value use map
package main

import (
	"fmt"
)

func main() {
	empSalary := make(map[string]int)
	fmt.Println(empSalary)

	//adding items into map
	empSalary["gaurav"] = 1000
	empSalary["danny"] = 1500
	fmt.Println(empSalary)
}
