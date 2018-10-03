package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main(){
	r := strings.NewReader("Some text string coming as the input\n")

	buf := make([]byte,40)
	if _, err := io.ReadAtLeast(r,buf,4); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Buffer: %s\n",buf)

	// Buffer smaller than read size.
	// The error is EOF only if no bytes were read.
	shortBuffer := make([]byte,4)
	if _, err := io.ReadAtLeast(r,shortBuffer,4); err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("Short Buffer: %s\n",shortBuffer)

	// Minimal read size is bigger than io.Reader stream.
	longBuf:= make([]byte,3)
	if _, err := io.ReadAtLeast(r,longBuf,65); err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("")

}
