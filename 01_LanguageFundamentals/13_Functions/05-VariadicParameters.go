package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {

	fmt.Println(sf) // let see what the sf looks like
	fmt.Printf("%T \n", sf)

	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
 * The final parameter in a function signature may have a type prefixed with ...
 * A function with such a parameter is called variadic and may be invoked with zero or more arguments of for that parameter
 * func (prefix string, values ... int)
 * func (a,b int, z float64, opt ..interface{})
 */
