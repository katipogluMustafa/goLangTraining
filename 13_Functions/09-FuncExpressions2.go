package main

import "fmt"

func makeGreeter() func() string { // makeGreeter Returns func() string
	return func() string {
		return "Hello, World!"
	}
}
func main() {
	greet := makeGreeter()
	fmt.Println(greet())

	fmt.Printf("%T\n", greet) // Return func() string
	/*
	 * It is considering the types as func along with the return
	 * In this case types is func() string
	 */
}

/*
 *  func             makeGreeter    ()          func() string
 *  	  Receiver - Identifier - Parameters - Return Type
 *      No Receiver				 No Parameters
 */
