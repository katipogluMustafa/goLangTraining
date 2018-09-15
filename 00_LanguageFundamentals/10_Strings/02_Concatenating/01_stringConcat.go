package main

import (
	"fmt"
	"strings"
)

// Inefficient Way of Concatenating
func join(strs ...string) string {
	var result string

	for _, str := range strs {
		result += str
	}
	return result
}

//Efficient Way of Concatenating
func concatenate(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}

	return builder.String()
}

func main() {
	fmt.Println(join("mustafa", " ", "katipoğlu"))
	fmt.Println(concatenate("mustafa", " ", "katipoğlu"))
}

/*
 * While join() method works perfectly fine for a simple program, it is a little inefficient.
 * Every time we call ret += str a brand new string needs to be allocated in memory.
 * This happens because strings in Go are immutable, so if we want to change a string or add contents to it we need to create an entirely new string.
 * To avoid creating new strings as we build our final string, we can now use the strings.Builder type along with its WriteString method.
 */
