package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "Mustafa Katipoğlu"
	h := sha256.New()
	h.Write([]byte(str))

	fmt.Println(str)
	fmt.Printf("Sha256 : %x", h.Sum(nil))
}
