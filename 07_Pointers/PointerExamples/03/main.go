package main

// Number of elements to grow each stack frame.
// Run with 10 and then with 1024
const size = 10

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

/*
 * stackCopy recursively runs increasing the size of the stack.
 */
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++

	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
