package main

import (
	"fmt"
	"strings"
)

func countWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}

func main() {
	str := "CountWords counts the number of words in the specified string and returns the count."
	fmt.Println(countWords(str))
}
