package main

import "fmt"

func zero1(z int) {
	z = 0
}
func zero2(z *int) {
	*z = 0
}

func main() {
	x := 5

	zero1(x)
	fmt.Println(x) // x is still 5

	zero2(&x)
	fmt.Println(x) // x is 0
}
