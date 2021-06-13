package main

import "fmt"

func main() {
	var i int8 = 1
	var j float32 = 50000000
	fmt.Println(i + int8(j))
	var k = 10.2
	fmt.Println(float64(k))
}
