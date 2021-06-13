package main

import (
	"fmt"
)

func main() {
	const a, b = 90, "gaurav"
	//	a = 30
	//	var a, b, c bool = false, true, true
	fmt.Println(a, b)

	const d = "gaurav" //must be inilitize value
	var c = d
	fmt.Println(c)

	//const d = math.Sqrt(4) //must be constant value at the time of compile
	fmt.Printf("type of e is %T, size of e is %d, size of e is %d", c, c, d)
	fmt.Println()

}
