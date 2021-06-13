package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b = 10.1, 20.1
	b, c := 30.1, 40.1
	fmt.Println(a, b, c, math.Min(a, b))
	var d bool
	d = true
	fmt.Println(d)
}
