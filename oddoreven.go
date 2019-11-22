package main

import "fmt"

func main() {
	checkNumber()
}

//check odd or even number
func checkNumber() {
	fmt.Print("Enter number : ")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println(n, "is an Even number")
	} else {
		fmt.Println(n, "is an Odd number")
	}
}
