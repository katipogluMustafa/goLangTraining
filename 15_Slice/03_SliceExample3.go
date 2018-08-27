package main

import "fmt"

func main() {

	greeting1 := []string{
		"Selamu Aleykum",
		"Günaydın!",
		"Good morning!",
		"Dias!",
		"Ohayo!",
	}

	fmt.Println("[1:2] ", greeting1[0:1])
	fmt.Println("[:2] ", greeting1[:2])
	fmt.Println("[3:] ", greeting1[3:])
	fmt.Println("[:] ", greeting1[:])

}
