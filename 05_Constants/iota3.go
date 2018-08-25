package main

import "fmt"

const (
	D = iota // 0
	E        // 1
	F        // 2
)
const (
	G = iota // 0
	H        // 1
	J        // 2
)

func main() {
	fmt.Println(D, E, F)
	fmt.Println(G, H, J)
}

/*

	As soon as we opened new const, it starts from scratch

*/
