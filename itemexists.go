package main

import (
	"fmt"
	"reflect"
)

//check if any item exists or not
func main() {
	var strSlice = []string{"ozombo", "twilio", "unicodedeveloper", "amycruz", "codebeast"}
	fmt.Println(itemExists(strSlice, "ozombo"))
	fmt.Println(itemExists(strSlice, "ichtrojan"))
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true //returned because ozombo is part of strSlice
		}
	}

	return false //returned because ichtrojan is missing from strSlice
}
