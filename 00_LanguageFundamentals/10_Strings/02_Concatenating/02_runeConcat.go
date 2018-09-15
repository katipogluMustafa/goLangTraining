package main

import (
	"fmt"
	"strings"
)

/*
 * You can use the WriteRune and WriteByte methods to add single characters to your string as you build it.
 */
func joinRunes(runes ...rune) string {
	var sb strings.Builder

	for _, r := range runes {
		sb.WriteRune(r)
	}
	return sb.String()
}

func main() {
	fmt.Println(joinRunes('a', 'b', 'c', '1', '😁', '😃'))
}
