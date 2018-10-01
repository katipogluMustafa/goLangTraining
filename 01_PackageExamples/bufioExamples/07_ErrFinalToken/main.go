package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main(){
	input := "abcdefghijklmmnoprst end sadasdsadasd"

	myScanner := bufio.NewScanner(strings.NewReader(input))
	myScanner.Split(split)
	for myScanner.Scan(){
		fmt.Println(myScanner.Text())
	}

	if myScanner.Err() != nil {
		fmt.Printf("Error: %s\n", myScanner.Err())
	}

}

func split(data []byte, atEOF bool)(advance int, token []byte, err error){
	advance, token, err = bufio.ScanWords(data,atEOF)

	if err == nil && token != nil && bytes.Equal(token, []byte{'e','n','d'}) {
		return 0 , []byte{'E','N','D'}, bufio.ErrFinalToken
	}

	return
}


/*
 * Both io.EOF and ErrFinalToken aren’t considered to be “true” errors — Err method will return nil if any of these two caused scanner to stop.
 */

/*
 * Scanner offers an option to signal so-called final token.
 * It’s a special token which doesn’t break the loop (Scan still returns true) but subsequent calls to Scan will stop immediately
 */
