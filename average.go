package main

import "fmt"

//calulate an average of a range of numbers entered
func main() {
	var num [100]int
	var temp, sum, avg int
	// chose number of elemets to work with
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&temp)
	for i := 0; i < temp; i++ {
		//enter numbers to calculate avergae from
		fmt.Print("Enter the number : ")
		fmt.Scanln(&num[i])
		sum += num[i]
	}

	avg = sum / temp
	fmt.Printf("The Average of entered %d number(s) is %d", temp, avg)
}
