package main

import (
	"fmt"
	"strconv"
)

func main() {
	// To Int
	i, _ := strconv.Atoi("-42")
	fmt.Printf("value: %v\ttype:%T\n", i, i)

	// To Bool
	b, _ := strconv.ParseBool("true")
	fmt.Printf("value: %v\ttype:%T\n", b, b)

	// To Float
	f, _ := strconv.ParseFloat("3.1465", 64)
	fmt.Printf("value: %v\ttype:%T\n", f, f)

	// To Int32
	int, _ := strconv.ParseInt("-45", 10, 32)
	in := int32(int)
	fmt.Printf("value: %v\ttype:%T\n", in, in)

	// To Uint
	u, _ := strconv.ParseUint("55", 10, 64)
	fmt.Printf("value: %v\ttype:%T\n", u, u)
}
