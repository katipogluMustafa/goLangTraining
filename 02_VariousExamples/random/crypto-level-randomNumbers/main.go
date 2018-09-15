package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	key := [20]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}
	fmt.Println(key)
}

/*
 * math/rand is much faster for applications that don’t need crypto-level or security-related random data generation.
 * crypto.rand is suited for secure and crypto-ready usage, but it’s slower.
 */
