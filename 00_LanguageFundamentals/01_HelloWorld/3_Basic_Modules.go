package main

import (
	"fmt"
	"math"
	"strings"
)

func main(){
	// Math Module
	fmt.Println( "Floor the 2.57845 to ", math.Floor(2.57845) )
	fmt.Println( "Absolute value of -255 is ", math.Abs(-255) )
	fmt.Println("Square Root of the 81 is ", math.Sqrt(81))
	fmt.Println("Ceil 2.57 = ", math.Ceil(2.57))

	// Strings Module
	fmt.Println( strings.Split("Lets_Just_Split_This", "_") )
	fmt.Println( "Comparison of 'Mustafa' and 'mustafa' is ", strings.Compare("Mustafa", "mustafa"))
	fmt.Printf("'afa' string is at the %d index of the Mustafa Katipoglu string",strings.Index("Mustafa Katipoglu", "afa"))
}