package main

import "fmt"

func main() {
	repeat("Github", 5)
}

func repeat(word string, reps int) {
	for i := 0; i < reps; i++ {
		fmt.Print(word)
	}
}
