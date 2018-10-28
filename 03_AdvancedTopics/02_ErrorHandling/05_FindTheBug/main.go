// Sample program to show see if the class can find the bug.

package main

import "log"

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
	if _, err = fail(); err != nil {
		log.Fatal("Why did this code failed ?")
	}

	log.Println("No Error")

}
