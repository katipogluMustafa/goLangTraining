package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* To Int */
	i := 23

	//Using strconv package
	t := strconv.Itoa(i)
	fmt.Printf("using strconv Package:\n value: %v, type: %T\n", t, t)

	//using fmt package
	s := fmt.Sprintf("%d", i)
	fmt.Printf("using fmt Package:\n value: %v, type: %T\n", s, s)

	/* To Other types */
	fmt.Println()

	// Bool to string
	strB := strconv.FormatBool(true)
	fmt.Printf("value: %v\ttype:%T\n", strB, strB)

	// Float to string
	strF := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("value: %v\ttype:%T\n", strF, strF)

	// Int to String
	strI := strconv.FormatInt(-42, 16)
	fmt.Printf("value: %v\ttype:%T\n", strI, strI)

	// Uint to String
	strU := strconv.FormatUint(42, 16)
	fmt.Printf("value: %v\ttype:%T\n", strU, strU)
}

/*
 * fmt.Sprintf operates pretty much identically to fmt.Printf except instead of printing out the resulting string to standard output it instead returns it as a string.
 */
/*
 * Limit your Sprintf usage
 * fmt.Sprintf should typically be reserved for creating strings with embedded values. There are a few reasons for this, but the most prominent one is that fmt.Sprintf doesn’t do any type checking, so you are unlikely to catch any bugs until you actually run your code.
 * fmt.Sprintf is also slower than most functions you would typically use in the strconv package, though if I am being honest the speed difference is so minimal it isn’t typically worth considering.
 */
