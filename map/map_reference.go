package main

import "fmt"

func main() {
	/*	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}*/

	employeeSalaryArray := [3]int{
		0: 12000,
		1: 15000,
		2: 9000,
	}
	/*	fmt.Println("Original map employee salary", employeeSalary)
		modified := employeeSalary
		modified["mike"] = 18000
		modified["gaurav"] = 18000
		fmt.Println("Employee map update salary changed", employeeSalary)
	*/
	//array and map differece
	fmt.Println("Original employee salary array ", employeeSalaryArray)
	modifiedArray := employeeSalaryArray
	modifiedArray[0] = 18000
	fmt.Println("Employee salary array changed", modifiedArray)
	fmt.Println("Employee salary array changed", employeeSalaryArray)

}
