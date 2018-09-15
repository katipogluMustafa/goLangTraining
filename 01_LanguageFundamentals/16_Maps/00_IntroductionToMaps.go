package main

import "fmt"

func main() {
	// Create an empty map using
	// make(map[key-type]val-type)
	m := make(map[string]int)

	//Set key/value pairs using
	// name[key] = val
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	// use name[key to get a value for a key
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	//  use len() to get length
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Println("ok?:", ok)
	/*
	 * The optional second return value when getting a value from a map indicates if the key was present in the map.
	 * This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	 * Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	 */

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, v)

	n := map[string]int{"foo": 1, "edu": 2}
	fmt.Println("map: ", n)
}

/*
 * Maps are for key - value association
 * Maps are unordered and Maps are a reference type
 * Maps are build on top of a hash table
 */
