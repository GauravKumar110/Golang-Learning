package main

import (
	"fmt"
)

type employee struct {
	salary     int
	country    string
	department string
}

func main() {
	emp1 := employee{
		salary:     12000,
		country:    "USA",
		department: "technolgy",
	}
	emp2 := employee{
		salary:     14000,
		country:    "Canada",
		department: "HR",
	}
	emp3 := employee{
		salary:     13000,
		country:    "India",
		department: "Business",
	}
	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s  Department: %s\n", name, info.salary, info.country, info.department)
	}

	fmt.Println("Length of employee : ", len(employeeInfo), cap(employeeInfo))

}
