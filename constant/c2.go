package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var defaultName = "gaurav"

	type myString string

	var cName myString = "kumar"

	//	cName = defaultName
	fmt.Println(cName)
	fmt.Printf("type of e is %T, size of e is %d, size of e is %d", defaultName, defaultName, unsafe.Sizeof(defaultName))
	fmt.Println()
	fmt.Printf("type of e is %T, size of e is %d, size of e is %d", cName, string(cName), unsafe.Sizeof(cName))
	fmt.Println()

}
