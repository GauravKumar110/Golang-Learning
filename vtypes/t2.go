package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int8 = -1
	var b int64 = 34333
	var c uint16 = 101

	fmt.Println(a, b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Println()
	fmt.Printf("type of b is %T, size of b is %d", b, unsafe.Sizeof(b))
	fmt.Println()
	fmt.Printf("type of c is %T, size of c is %d, size of c is %d", c, c, unsafe.Sizeof(c))
	fmt.Println()

	var d float32 = 5.67
	var e = 6.57

	fmt.Printf("type of d is %T, size of d is %d", d, unsafe.Sizeof(d))
	fmt.Println()
	fmt.Printf("type of e is %T, size of e is %d", e, unsafe.Sizeof(e))
	fmt.Println()
	var f = "gaurav"
	fmt.Printf("type of e is %T, size of e is %d, size of e is %d", f, f, unsafe.Sizeof(f))
	fmt.Println()

}
