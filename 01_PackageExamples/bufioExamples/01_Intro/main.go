package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main(){

	s := "This is a string \n which has some line breaks \n but at the end it is a string \n isn't it ?"

	// Use a wrapper to make string a Reader then create a new scanner using bufio package
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		// Scanner by default return each line of the content, but the behaviour can be changed using Split() method.
		line := scanner.Text()
		fmt.Println(line)
	}
}