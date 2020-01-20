package main

import (
	"fmt"
	"math"
)

func main() {
	mathmethods()
	fahrenheit()
	feetconversion()
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

// feet to metre conversion
func feetconversion() {
	// define feet for use as feet to convert to metres
	var feet float32
	fmt.Println("Enter feet length to convert")
	fmt.Scan(&feet)
	var feetvalue float32
	feetvalue = 3.28
	result := (feet / feetvalue)
	// fmt.Println(feet, "feet is", int(result), "metres")
	fmt.Println(feet, "feet is", result, "metres")
}

// farenheit to celcius conversion
func fahrenheit() {
	// define farenheit for use as the degress in farenheit to convert
	var fahrenheit int
	// ask for the farenheit value to be converted
	fmt.Println("Enter farenheit degrees to convert")
	// scan the inputed integer
	fmt.Scan(&fahrenheit)
	// perform farenheit to celcius calculation
	celcius := (fahrenheit - 32) * 5 / 9
	// print the result
	fmt.Println(fahrenheit, "in celcius is", celcius, "degree celcius")
}
