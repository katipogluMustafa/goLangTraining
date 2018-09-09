package main

import (
	"fmt"
	"math/rand"
)

// returns int, min <= int < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := range bytes {
		bytes[i] = byte(randomInt(97, 123))
	}
	return string(bytes)
}

func main() {
	str := randomString(200)
	fmt.Println(str)
}
