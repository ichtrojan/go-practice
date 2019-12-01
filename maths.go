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

	num1 := 10
	num2 := 2
	fmt.Println("The square root of 16 is", math.Sqrt(16)) //square root
	fmt.Println(num1, "+", num2, "=", num1+num2)           // Addition
	fmt.Println(num1, "*", num2, "=", num1*num2)           // Multiplication
	fmt.Println(num1, "-", num2, "=", num1-num2)           // Subtraction
	fmt.Println(num1, "/", num2, "=", num1/num2)           // Division
	fmt.Println(num1, "%", num2, "=", num1%num2)           // Remainder

	u := (10 + 10) * 5
	fmt.Println("10 + 10 * 5 = ", u)
}
