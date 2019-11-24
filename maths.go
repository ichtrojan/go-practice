package main

import (
	"fmt"
	"math"
)

func main() {
	mathmethods()
}

func mathmethods() {
	a := math.Min(100, 20)
	fmt.Println("Minimum number between 100 and 20 is", a)

	b := math.Max(100, 20)
	fmt.Println("Maximum number between 100 and 20 is", b)

	fmt.Println("The square root of 16 is", math.Sqrt(16)) //square root
	fmt.Println("1 + 2 = ", 1+2)                           // Addition
	fmt.Println("5 * 2 = ", 5*2)                           // Multiplication
	fmt.Println("10 - 2 = ", 10-2)                         // Subtraction
	fmt.Println("10 / 2 = ", 10/2)                         // Division
	fmt.Println("10 % 2 = ", 10%2)                         // Remainder

	u := (10 + 10) * 5
	fmt.Println("10 + 10 * 5 = ", u)
}
