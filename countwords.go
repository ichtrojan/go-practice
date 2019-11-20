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
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	strLine := "ozombo unicodedeveloper node ichtrojan codebeast vscode ozombo ichtrojan twitter codebeast ozombo"
	for index, element := range wordCount(strLine) {
		fmt.Println(index, "=>", element)
	}
}
