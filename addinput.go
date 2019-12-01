package main

import (
	"fmt"
)

func main() {
	firstNumber, secondNumber := getUserInput()
	sum := addTwoNumbers(firstNumber, secondNumber)
	fmt.Println(firstNumber, "+", secondNumber, "=", sum)
}

func addTwoNumbers(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func getUserInput() (int, int) {
	var firstNumber int
	var secondNumber int

	fmt.Println("Enter first number: ")
	fmt.Scan(&firstNumber)
	fmt.Println("Enter second number: ")
	fmt.Scan(&secondNumber)

	return firstNumber, secondNumber
}
