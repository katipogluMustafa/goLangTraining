// Sample program to show how to use error variables
// to help the caller determine the exact error being returned.

package main

import (
	"errors"
	"fmt"
)

/*
 * Using Err as the prefix for error variables is convention for errors.
 */
var (
	// ErrBadRequest is returned when there are problems with the request.
	ErrBadRequest = errors.New("Bad Request")

	// ErrPageMoved is returned when a 301/302 is returned.
	ErrPageMoved = errors.New("Page Moved")
)

func main() {
	if err := webCall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad request occured.")
			return
		case ErrPageMoved:
			fmt.Println("Sorry, this paged moved.")
			return
		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("end.")
}

func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}

	return ErrPageMoved
}

/*
 * I want this error variables which used in multiple places at the top of the source code file where they're used.
 */
