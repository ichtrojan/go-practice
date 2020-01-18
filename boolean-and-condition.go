package main

import "fmt"

func main() {
	subject := "Name" //Change name to your name of choice
	single := true

	if single {
		fmt.Println(subject, "is Single")
	} else {
		fmt.Println(subject, "is not single")
	}

	employed := false

	if employed {
		fmt.Println(subject, " is Employed")
	} else {
		fmt.Println(subject, "is Unemployed and Jobless at the same time")
	}

	// and statement
	if employed && single {
		fmt.Println(subject, "is single and employed")
	} else {
		fmt.Println(subject, "is single and jobless")
	}

	// or statement
	if employed || single {
		fmt.Println(subject, "is single or employed")
	} else {
		fmt.Println(subject, "is not single and is jobless")
	}
}
