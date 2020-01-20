package main

import (
	"fmt"
)

func main() {
	// define slice1 as four numbers
	slice1 := []int{1, 2, 3, 4}
	// join slice1 into slice2
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println("Adding slice1 into slice2 to give", slice2)

	// define slice 3 as three numbers
	slice3 := []int{1, 2, 3}
	// make slice4 as two numbers
	slice4 := make([]int, 2)
	// copy slice3 into slice4
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
