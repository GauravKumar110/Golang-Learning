//assign one array into second array

package main

import "fmt"

func main() {

	a := [...]int8{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	b := a
	fmt.Println(b)
	var c [8]int8 //size and type both are same as previous array
	c = a
	fmt.Println(c)

	//print array
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	//using range
	var total int8

	for key, value := range c {
		fmt.Println("key = ", key, " and value is = ", value)
		total += value
	}

	fmt.Println("Total is : ", total)

}
