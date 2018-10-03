package ex2

import (
	"fmt"
	"io"
	"strings"
)

type Example struct {
	ReaderLength int
	minimum 	 int
	text 		 string
}

func main(){

	examples := []Example{
		{10, 5, "Go is a programming language"},
		{10,5, "Golang"},
		{10, 0, "Golang2"},
		{20,15,"Mustafa, code, Go code, Golang"},
	}

	for _, ex := range examples{
		rd := strings.NewReader(ex.text)

		buffer := make([]byte, ex.ReaderLength)

		bytesRead, err := io.ReadAtLeast(rd,buffer,ex.minimum)

		fmt.Printf("\nLen(buffer): %v, Min:%d BytesRead: %d, Error: %v , %s \n", ex.ReaderLength, ex.minimum, bytesRead, err ,ex.text)

	}


}

// ReadAtLeast reads from r into buf until it has read at least min bytes.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading fewer than min bytes,
// ReadAtLeast returns ErrUnexpectedEOF.
// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.