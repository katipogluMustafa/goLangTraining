package main

import "fmt"

func main() {

	// Create map using Make
	var myGreeting = make(map[string]string)
	myGreeting["Mustafa"] = "Selamu Aleyküm"
	myGreeting["Todd"] = "Good morning"

	fmt.Println(myGreeting)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// Short way of creating map using make
	myGreeting1 := make(map[string]string)
	myGreeting1["Yasin"] = "Selamu Aleyküm"
	myGreeting1["Todd"] = "Good morning"

	fmt.Println(myGreeting1)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// Create map using composite literal
	var myGreeting2 = map[string]string{}
	myGreeting2["Fatih"] = "Selamu Aleyküm"
	myGreeting2["Tim"] = "Good morning"

	fmt.Println(myGreeting2)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// Shorthand Composite literal
	myGreeting3 := map[string]string{
		"Yusuf":   "Selamu Aleyküm",
		"Mustafa": "Aleyküm Selam",
	}

	myGreeting3["Mahmut"] = "Aleyküm Selam"

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// Length of the map
	fmt.Println("Length: ", len(myGreeting3))

	fmt.Println(myGreeting3)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// Update an Entry of Map

	myGreeting3["Mahmut"] = "Ve Aleyküm Selam"

	fmt.Println(myGreeting3)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// Delete an Entry

	delete(myGreeting2, "Tim")

	fmt.Println(myGreeting2)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	// nil map
	var myGreeting4 map[string]string
	// myGreeting3["Ahmet"] = "Selamu Aleyküm"
	// nil map, if you try to add, you'r gonna get panic
	fmt.Println(myGreeting4 == nil)

}

/* From the blog at golang.org:

 * This variable m is a map of string keys to int values:

 * var m map[string]int
 * Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map.
 * A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:

 * m = make(map[string]int)
 * The make function allocates and initializes a hash map data structure and returns a map value that points to it.
 * The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself.

 */
