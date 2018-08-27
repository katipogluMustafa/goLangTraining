package main

import "fmt"

func main() {
	var x [223]string
	fmt.Println(len(x))
	fmt.Println(x[0])

	for i := 33; i < 256; i++ {
		x[i-33] = string(i)
	}

	for _, v := range x {
		fmt.Printf("%v - % T - %v\n", v, v, []byte(v))
	}
}
