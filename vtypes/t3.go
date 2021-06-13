package main

import "fmt"

func main() {
	var c1 = complex(7, 5)
	var c2 = 7 * 27i

	var cadd = c1 + c2
	var cmul = c1 * c2

	fmt.Println("sum := ", cadd)
	fmt.Println("product := ", cmul)

}
