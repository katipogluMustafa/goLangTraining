package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sort strings
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Sort ints
	ints := []int{8, 1, 6, 2, 5, 4, 8}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// is sorted?
	b := sort.IntsAreSorted(ints)
	fmt.Println("is sorted?", b)
}
