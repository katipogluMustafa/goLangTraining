package main

import "fmt"

func main() {
	s := "this is a string"
	fmt.Println(s)

	var b []byte
	fmt.Println(b)

	b = []byte(s)
	fmt.Println(b)

	for i := range b {
		fmt.Println(string(b[i]))
	}
	s = string(b)
	fmt.Println(b)
}
