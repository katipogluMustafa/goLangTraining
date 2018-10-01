package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func main(){
	input := "abcdefghijklmmn"
	scanner := bufio.NewScanner(strings.NewReader(input))

	/*
	 * atEOF parameter designed to pass information to split function that no more data will be available.
	 * It can happen either while reaching EOF or if read call returns an error.
	 * If any of these happens then scanner will never try to read anymore.
	 */
	split := func(data []byte, atEOF bool)(advance int, token []byte, err error){
		fmt.Printf("%t\t%d\t%s\n",atEOF,len(data),data)
		if atEOF {
			return 0,nil,errors.New("uu error!")
		}
		return 0,nil,nil
	}

	scanner.Split(split)

	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan(){
		fmt.Printf("%s\n",scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf("error: %s\n",scanner.Err())
	}
}
