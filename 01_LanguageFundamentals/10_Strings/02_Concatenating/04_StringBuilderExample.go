package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("Surprise!")
	fmt.Println(b.String())
}

/*
 * The string builder also implements the io.Writer interface
 * We can use functions like fmt.Fprintf along with the string builder.
 */
