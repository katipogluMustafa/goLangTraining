package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main(){
	input := "GoGoGoGo"

	scanner := bufio.NewScanner(strings.NewReader(input))

	split := func(data []byte, atEOf bool)(advance int, token []byte, err error){
		if bytes.Equal(data[:2], []byte{ 'G', 'o'}) {
			return 2, []byte{'G'}, nil
		}
		if atEOf {
			return 0,nil,io.EOF
		}

		return 0, nil , nil
	}

	scanner.Split(split)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
}
