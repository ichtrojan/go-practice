package main

import "fmt"

func main() {
	var choice int
	var firstNumber float32
	var secondNumber float32

	fmt.Println("Select an Operation")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("Enter choice(1/2/3/4): ")
	fmt.Scan(&choice)
	fmt.Println("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Println("Enter second number: ")
	fmt.Scan(&secondNumber)

	if choice == 1 {
		fmt.Println(firstNumber, "+", secondNumber, "=", firstNumber+secondNumber)
	} else if choice == 2 {
		fmt.Println(firstNumber, "-", secondNumber, "=", firstNumber-secondNumber)
	} else if choice == 3 {
		fmt.Println(firstNumber, "*", secondNumber, "=", firstNumber*secondNumber)
	} else if choice == 4 {
		fmt.Println(firstNumber, "/", secondNumber, "=", firstNumber/secondNumber)
	} else {
		fmt.Println("Invalid operation selected")
	}
}
