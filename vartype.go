package main

import (
	"fmt"
)

// determine type of variable
func main() {
	msg := "Hello"
	fmt.Printf("%v is %T\n", msg, msg)
	num := 32
	fmt.Printf("%v is %T\n", num, num)
	val := 1 == 1 //check boolean
	fmt.Printf("Result is is %T\n", val)
}
