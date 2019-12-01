package main

import (
	"fmt"
	"reflect"
)

//check if any item exists or not
func main() {
	var strSlice = []string{"banana", "apple", "peach", "managa", "dates"}
	fmt.Println(itemExists(strSlice, "apple"))
	fmt.Println(itemExists(strSlice, "orange"))
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true //returned because apple is part of strSlice
		}
	}

	return false //returned because orange is missing from strSlice
}
