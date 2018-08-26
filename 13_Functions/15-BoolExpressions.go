package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(!true) // ! is the not operator
	fmt.Println(5 > 7) // Evaluates to a bool
	fmt.Println(4 > math.Pi && 4 < 2 * math.Pi) // Evaluates to a bool
}

/*
 * Bool expressions that evaluate down to a value of type bool
 * Statements perform action , expression produces a value
 */
