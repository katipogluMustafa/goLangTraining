// Sample program to show how the default error type is implemented.

package main

import "fmt"

/* http://golang.org/pkg/builtin/#error */
type error interface {
	Error() string
}

/* http://golang.org/src/pkg/errors/errors.go */
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New return an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

func main() {
	if err := webCall(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Life is perfect.")
}

func webCall() error {
	return New("Bad Request")
}
