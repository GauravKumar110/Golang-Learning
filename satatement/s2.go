package main

import (
	"fmt"
)

func main() {
	//num := 11
	if var num = 5, num%2 == 0 { //checks if number is even
		fmt.Println("The number", num, "is even")
		return
	} 
	else {
		fmt.Println("The number", num, "is odd")
	}
}
