package main

import (
	"bufio"
	"fmt"
	"strings"
)

/*
 * By default maximum length of buffer which is used underneath is 64 * 1024 bytes.
 * It means that found token cannot be longer than this limit (source code)
 */

 // See The Error
func main(){
	// Repeat returns a new string consisting of count copies of the string s.
	input := strings.Repeat("x",bufio.MaxScanTokenSize)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

// Program prints bufio.Scanner: token too long.