package main

import (
	"fmt"
	"sort"
)

func main() {
	myInts := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	myStrings := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(myInts)
	// sort myInts
	sort.Ints(myInts)
	fmt.Println(myInts)

	fmt.Println(myStrings)
	// sort myStrings
	sort.Strings(myStrings)
	fmt.Println(myStrings)

}
