package main

import "fmt"

func main() {
	var x [58]string
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}

/*
 * If we want to store the numbers between 65 to 122
 * 122 - 65 = 57 but since we include both boundaries
 * we need 58 spot, so the length is 58
 * just one boundary included, length will be difference
 * no boundary included, length will be difference - 1
 * both boundary included, length will be difference + 1
 */
