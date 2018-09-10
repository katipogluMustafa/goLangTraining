package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "Mustafa Katipoğlu"

	hash := sha1.New()
	hash.Write([]byte(str))

	bs := hash.Sum(nil)

	fmt.Println(str)
	fmt.Printf("Sha1 : %x", bs)
}
