/* How the for range has pointer semantics. */
package main

import "fmt"

func main() {
	// Using the pointer semantic form of the for range.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i])
	}
}

/*
 * When we use pointer semantics form of for range, we start ranging over original slice of friends
 * since the initial length of friends is 5 for range tries to range over the slice 5 times but the change of the friends array caused the array to shrink and now we can't access to any more than 2 index.
 */
