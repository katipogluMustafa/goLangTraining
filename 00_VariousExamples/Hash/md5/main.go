package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "Mustafa Katipoğlu"

	hash := md5.New()
	hash.Write([]byte(str))

	bs := hash.Sum(nil)

	fmt.Println(str)
	fmt.Printf("md5 : %x", bs)
}
