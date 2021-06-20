package main

import (
	"fmt"
)

func subtactOne(numbers [3]int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}
func main() {
	nos := [3]int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside
}
