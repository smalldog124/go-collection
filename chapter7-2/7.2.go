package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5, 'k'}
	fmt.Println(slice2)
	fmt.Println("Slice with length and capacity")
	fmt.Printf("slice: length %v, capacity %v, %v\n", len(slice2), cap(slice2), slice)
}
