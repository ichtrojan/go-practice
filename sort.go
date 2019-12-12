package main

import (
	"fmt"
	"sort"
)

//sorts alphabets alphabetically and string incrementally
func main() {
	alph := []string{"z", "v", "a", "d", "g", "j"}
	sort.Strings(alph)
	fmt.Println("Alphabets:", alph)

	integer := []int{71, 902, 43}
	sort.Ints(integer)
	fmt.Println("Integer:   ", integer)

	s := sort.IntsAreSorted(integer)
	fmt.Println("Sorted?: ", s)
}
