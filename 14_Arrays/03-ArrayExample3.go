package main

import "fmt"

func main() {
	/* byte is alias for uint8 */
	var x [256]byte

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %b \n", v, v, v)
		if i > 50 {
			break
		}
	}

}
