//passing a slice into variadic function

package main

import (
	"fmt"
)

func find(num float32, n1 float32, nums ...float32) {
	fmt.Printf("type of nums is %T\n", nums)
	fmt.Println(nums)
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
	//nums := []int{89, 90, 95}
	find(89.44, 89.44)
}
