// Sample program to show see if the class can find the bug.

package main

import (
	"fmt"
	"log"
)

// customError is an empty struct.
type customError struct {
}

// customError implements error interface.
func (c *customError) Error() string {
	return "Find the bug."
}

func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {

	var err error
	fmt.Printf("Type of value stored in the interface: %T\n", err)
	if _, err = fail(); err != nil {
		fmt.Printf("Type of value stored in the interface: %T\n", err)
	}

	log.Println("No Error")

}
