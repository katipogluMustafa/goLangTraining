package main

import "fmt"

func main() {

	greeting := []string{
		"Selamu Aleykum",
		"Günaydın!",
		"Good morning!",
		"Dias!",
		"Ohayo!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
