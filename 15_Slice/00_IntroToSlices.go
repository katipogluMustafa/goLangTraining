package main

import "fmt"

func main() {
	/*
	 * Slice is a reference type.
	 */
	// Example 1
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	/*
	 * A Slice is a descriptor for a contiguous segment of an underlying array
	 * and provides access to a numbered sequence of elements from that array.
	 */
	// Example 2
	mySecondSlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySecondSlice)
	fmt.Println(mySecondSlice[2:4]) /* get the part starting from 2 to 4 but 4 is not included */
	fmt.Println(mySecondSlice[2])   // Index access

	fmt.Println("myString"[2])         // Index access
	fmt.Println(string("myString"[2])) // Index access
	/*
	 * You can do slicing on a string because a string is a slice of bytes.
	 * 	A string is made up of runes
	 * 	A rune is a unicode code point
	 * 	A unicode code point is 1 to 4 bytes.
	 * Strings are made up of runes, runes are made up of bytes, so strings are made up of bytes.
	 * A string is a bunch of bytes; a slice of bytes.
	 */
}

/*
 * The value of an uninitialized slice is nil
 */

/*
 * Like arrays, slices are indexable and have a length.
 * The length of a slice s can be discovered by the built-in function len
 * Unlike Arrays, slices are dynamic
 * 	Their length may change during execution
 */
