package main

import "fmt"

//get first and last element of a slice and remove last
func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", intSlice)

	last := intSlice[len(intSlice)-1]
	fmt.Printf("Last element: %v\n", last)

	first := intSlice[:1]
	fmt.Printf("First element: %d\n", first)

	remove := intSlice[:len(intSlice)-1]
	fmt.Printf("Remove Last: %v\n", remove)
}
