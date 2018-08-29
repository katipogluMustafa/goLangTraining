package main

import (
	"fmt"
	"io"
	"crypto/rand"
)

func main(){
	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)

	// Line 12 makes the nonce a slice that points to underlying array.
}