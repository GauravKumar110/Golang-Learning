package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[1:3]
	fmt.Println("array before slice: ", arr)
	fmt.Println(" slice: ", slice)

	for i, v := range slice {
		slice[i]++
		fmt.Println("key = ", i, "value = ", v, "slice[i] = ", slice[i])
	}

	fmt.Println("after after", arr)
}
