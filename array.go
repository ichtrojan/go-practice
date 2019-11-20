package main

import "fmt"

func main() {
	// Declare an array to hold 5 integers
	numbers := [5]int{1, 2, 3, 4, 5}

	// Print a single element
	fmt.Println(numbers[3])

	// Loop through the array to print every value
	for _, number := range numbers {
		fmt.Println(number)
	}
}
