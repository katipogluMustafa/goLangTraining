/* How the for range has value semantics. */

package main

import "fmt"

func main() {

	// Using the value semantic form of the for range.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}
}

/*
 * Since the values semantic form of for range
 * creates its own copy of friends before ranging over it
 * even if we change the original friends slice
 * change doesn't effect us since we are ranging over our own copy
 */
