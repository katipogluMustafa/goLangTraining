package main

import "fmt"

func main() {
	var x [256]int

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %b \n", v, v, v)
		if i > 50 {
			break
		}
	}

}

/*
 * Between 0 to 256[not included] there are 256 spots
 * If we want to store 256 item in a array
 * we need to make the length of the array 256
 * and store elements from index 0 to 255
 * we don't store anything on index 256
 */
