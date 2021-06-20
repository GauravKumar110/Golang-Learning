package main

import (
	"fmt"
)

type Country struct {
	country string
}

type Address struct {
	city    string
	state   string
	country Country
}

type Salary struct {
	empSalary int
}

type Person struct {
	name    string
	age     int
	address Address
	Salary
}

func main() {
	p := Person{
		name: "Naveen",
		age:  50,
		address: Address{
			country: Country{
				"india",
			},
			city:  "Chicago",
			state: "Illinois",
		},
		Salary: Salary{
			10000,
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Salary:", p.empSalary)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
	fmt.Println("State:", p.address.country.country)
}
