package main

import (
	"fmt"
)

func main() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	veggies = append(veggies, fruits...)
	fmt.Println("food:", veggies)
}
