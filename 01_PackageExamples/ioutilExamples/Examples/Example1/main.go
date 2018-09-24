package main

import (
	"bytes"
	"fmt"
	"os"
)

// The io.Reader and io.Writer interfaces allow you to compose all of these different bits together.

/*
 * Sample program to show how different functions from the standard library use the io.Writer interface.
 */

func main() {

	// Create a buffer value and write a string to the buffer.

	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// Use Fprintf to concatenate a string to the buffer.
	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)

}
