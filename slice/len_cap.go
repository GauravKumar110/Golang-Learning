package main

import (
	"fmt"
)

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Println(fruitslice)
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
}
