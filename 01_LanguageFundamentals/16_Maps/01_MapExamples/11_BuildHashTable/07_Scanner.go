package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	const input = "asd asdsadas  dsada sdasd"

	 scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
