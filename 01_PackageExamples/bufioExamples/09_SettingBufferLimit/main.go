/* The buffer limit can be set with method Buffer which allows also to pass custom buffer.*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main(){
	buf := make([]byte, 10)
	// Repeat returns a new string consisting of count copies of the string s.
	input := strings.Repeat("x",20)
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set max limit of buffer.
	scanner.Buffer(buf,20)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

