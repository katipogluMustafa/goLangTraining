package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter digits to sum separated by space: ")

	str, err := getline()

	if err == nil {
		fmt.Println("Result: ", sumNumbers(str))
	}
}

func getline() (string, error) {
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

func sumNumbers(str string) float64 {
	var sum float64

	// Fields splits the string str around each instance of one or more consecutive white space
	// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
	// empty slice if str contains only white space.
	fmt.Println(strings.Fields(str))

	for _, v := range strings.Fields(str) {

		// ParseFloat converts the string s to a floating-point number
		i, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Println(err)
		} else {
			sum += i
		}

	}

	return sum
}
