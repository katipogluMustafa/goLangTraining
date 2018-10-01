package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func main(){
	input := "abcdefghijklmmnoprst"
	scanner := bufio.NewScanner(strings.NewReader(input))

	split := func(data []byte, atEOF bool)(advance int, token []byte, err error){
		return 0, nil, errors.New("No no no no, it's error, sorry :( ")
	}
	scanner.Split(split)

	for scanner.Scan() {
		fmt.Printf("%s\n",scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Printf("error : %s\n", scanner.Err())
	}

}

/* If split function returns an error then scanner stops */