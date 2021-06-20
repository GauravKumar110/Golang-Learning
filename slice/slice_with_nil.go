package main

import "fmt"

func main() {

	var name []string

	if name == nil {
		fmt.Println("len of name : ", len(name), "cap of name : ", cap(name))
		name = append(name, "a", "b", "c")
		fmt.Println("new names", name)
		fmt.Println("len of name : ", len(name), "cap of name : ", cap(name))
	}
}
