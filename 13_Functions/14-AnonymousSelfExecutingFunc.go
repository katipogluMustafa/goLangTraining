package main

import "fmt"

func main() {
	/* Example 1 */
	func() {
		fmt.Println("I'm driving!")
	}()

	/* Example 2 */
	fmt.Println(
		func() int {
			return 5
		}())

	/* Example 3 */
	func(str string) {
		fmt.Println(str)
	}("This is a string")
}

/*
 * Anonymous self executing functions have no names
 * optionally take parameters, return somethings
 */
