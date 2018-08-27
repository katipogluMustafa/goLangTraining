package main

import "fmt"

func main() {

	// Example 1

	customerNumber := make([]int, 3) // 3 is length & capacity
	/*
	 * length - number of elements referred to by the slice
	 * capacity - number of elements in the underlying array
	 */
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	for _, v := range customerNumber {
		fmt.Println(v)
	}

	// Example 2

	greeting3 := make([]string, 3, 5)
	greeting3[0] = "Selamu Aleyküm"
	greeting3[1] = "Günaydın!"
	greeting3[2] = "Dias!"
	// greeting3[3] = "good morning" // Index Out of range Error!
	greeting3 = append(greeting3, "Good morning!")
	/*
	 * If we try to add greeting[3] = "good morning" here
	 * we get index out of range error
	 * since this slice's length is set to 3, if we want to add more elements we need to use append()
	 */

	fmt.Println(greeting3[0])

}

/*
 * If we think that our slice might grow,
 * we can set a capacity larger than length
 * This gives our slice room to grow without golang having to create a new underlying array every time our slice grows.
 * When the slice exceeds capacity, then a new underlying array will be created.
 * These arrays double in size each time they're created(2,4,8,16) up to a certain point, and then they scale in some smaller proportion.
 */
