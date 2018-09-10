package main

import (
	"fmt"
	"strings"
)

func joinedAndRevesed(strs ...string) (string, string) {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	joined := sb.String()

	//Reset The String builder
	sb.Reset()

	for i := len(strs) - 1; i >= 0; i-- {
		sb.WriteString(strs[i])
	}
	reverse := sb.String()

	return joined, reverse
}

func main() {
	joined, reversed := joinedAndRevesed("mustafa", "katipoğlu")
	fmt.Printf("Joined: %s\nReversed: %s\n", joined, reversed)
}
