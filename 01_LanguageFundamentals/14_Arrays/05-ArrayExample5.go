package main

import "fmt"

func main() {

	/*
	 * The expression on the right-hand side of the above statement is called an array literal.
	 * Note that we do not need to specify the type of the variable a as in var a [5]int
	 * because the compiler can automatically infer the type from the expression on the right hand side.
	 */
	var b = [5]int{2, 4, 6, 8, 10} // Initializing an array using an array literal
	fmt.Println(b)

	c := [5]int{2, 4, 6, 8, 10} // Short hand declaration
	fmt.Println(c)

	/*
	 * Let the compiler count the number of elements for you
	 */
	a := [...]int{3, 5, 7, 9, 11, 13, 17} // Let the Go compiler infer the length of the array
	fmt.Println(a)

}
