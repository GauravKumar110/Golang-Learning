//Accessing individual bytes of a string

package main

import (
	"fmt"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i])
		fmt.Printf("%c \n", s[i])
	}

	fmt.Printf("%x %x", "a", "A")
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}
