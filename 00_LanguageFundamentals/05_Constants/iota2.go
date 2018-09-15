package main

import "fmt"

const (
	X     = iota // 0
	Y            // 1
	Z            // 2
	think = "hi"
	W     // won't be 3
)

func main() {
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
	fmt.Println(W)
	fmt.Println(think)
}

/*

	We don't have to say the variable equal to iota as long as there is no other types of constant in the middle.

*/
