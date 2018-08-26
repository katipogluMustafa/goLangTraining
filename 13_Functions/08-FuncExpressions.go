package main

import "fmt"

func main() {
	/* Here is a Func Expression */
	greeting3 := func() {
		fmt.Println("Hello, World!")
	}

	greeting3()
	fmt.Printf("%T\n", greeting3) // Type of greeting3 is a func()
}

/*
 * Func expression is the only way we could have func inside another func
 */
