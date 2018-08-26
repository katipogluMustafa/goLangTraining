package main

import "fmt"

/*
 * Visit function takes two parameters
 * It takes a slice of int and a func(int)
 */
func visit(numbers []int, callback func(int)) {// Here callback is just identifier, we could name it anything
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	myFunc := func(n int) {
		fmt.Println(n)
	}
	visit([]int{1, 2, 3, 4}, myFunc)
}
/*
 * Callback is passing function as an argument
 */

