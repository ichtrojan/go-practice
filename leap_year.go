package main

import "fmt"

//check leap year
func checkLeapYear() {
	fmt.Print("Enter a year to check: ")
	var year int
	fmt.Scanln(&year)
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println("The year", year, "is a leap year")
			} else {
				fmt.Println("The year", year, "is not a leap year")
			}
		} else {
			fmt.Println("The year", year, "is a leap year")
		}
	} else {
		fmt.Println("The year", year, "is a not a leap year")
	}
}

func main() {
	checkLeapYear()
}
