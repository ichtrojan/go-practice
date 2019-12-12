package main

import (
	"fmt"
	"strings"
)

//count appreance of words in a string
func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	strLine := "I love writing code, I used to work with PHP but now I am loving GO more! Do you code in GO"
	for index, element := range wordCount(strLine) {
		fmt.Println(index, "=>", element)
	}
	fmt.Println("There are", len(strLine), "characters in the defined string.") //count lenght of string
}
