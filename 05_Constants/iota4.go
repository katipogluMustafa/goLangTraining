package main

import "fmt"

const (
	_ = iota      //let's throw the 0 away then use others
	K = iota * 10 // 1 * 10
	L = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(K)
	fmt.Println(L)
}
