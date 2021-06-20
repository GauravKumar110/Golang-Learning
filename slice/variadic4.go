//Slice arguments vs Variadic arguments

package main

import (
	"fmt"
)

func find(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	fmt.Println("type of nums is", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find(89, []int{89, 90, 95})
	//find(45, []int{56, 67, 45, 90, 109})
	//find(78, []int{38, 56, 98})
	//find(87, []int{})
}