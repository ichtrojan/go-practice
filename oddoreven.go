package main

import "fmt"

//check odd or even number
func main() {
	fmt.Print("Enter number : ")
	var n int
	fmt.Scanln(&n)
	/*  Conditional Statement if .... else ........     */
	if n%2 == 0 {
		fmt.Println(n, "is an Even number")
	} else {
		fmt.Println(n, "is an Odd number")
	}
}
