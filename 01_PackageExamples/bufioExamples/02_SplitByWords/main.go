package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main(){

	s := "NewScanner returns a new Scanner to read from io.Reader. \n The split function defaults to ScanLines."

	scanner := bufio.NewScanner(strings.NewReader(s))

	// Change the split behaviour from splitting the lines to splitting words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		line := scanner.Text()
		fmt.Println(line)
	}

}
