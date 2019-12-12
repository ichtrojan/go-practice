package main

import "fmt"

func main() {
	single := true

	if single {
		fmt.Println("Trojan is Single")
	} else {
		fmt.Println("Trojan is not single")
	}

	employed := false

	if employed {
		fmt.Println("Kofo is Employed")
	} else {
		fmt.Println("Kofo is Unemployed and Jobless at the same time")
	}

	// and statement
	if employed && single {
		fmt.Println("Subject is single and employed")
	} else {
		fmt.Println("Subject is single and jobless")
	}

	// or statement
	if employed || single {
		fmt.Println("Subject is single or employed")
	} else {
		fmt.Println("Subject is not single and is jobless")
	}
}
