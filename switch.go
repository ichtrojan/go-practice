package main

import "fmt"

func main() {
	// Declare the desire
	desire := "money"

	switch desire {
	case "love":
		fmt.Println("❤️❤️❤️❤️❤️❤️❤️")

	case "money":
		fmt.Println("💰💰💰💰💰💰💰")

	default:
		fmt.Println("I cannot give you what you desire")
	}
}
