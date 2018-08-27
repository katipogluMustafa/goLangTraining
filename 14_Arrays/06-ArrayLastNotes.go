package main

import "fmt"

func main() {

	/* Array’s length is part of its type */
	/*
	 * The length of an array is part of its type.
	 * The array a[8]int and a[12]int are totally distinct types, and you cannot assign one to the other.
	 */
	//var a = [5]int{3, 5, 7, 9, 11}
	//var b [10]int = a  // Error, a and b are distinct types

	/* Arrays in Golang are value types */
	/*
	 * When you assign an array to a new variable or pass an array to a function, the entire array is copied.
	 * So if you make any changes to this copied array, the original array won’t be affected, and will remain unchanged.
	 */
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)

	/* Iterating over an array using range */

	// Example 1
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}

	// Example 2
	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	var sum float64

	for _, value := range a {
		sum = sum + value
	}
	fmt.Println("Sum: ", sum)

	/* Multidimensional arrays in Golang */

	// Example 1
	d := [2][2]int{
		{3, 5},
		{7, 9}, // This comma is necessary /* !!!!!!!!!!!!!!!!!!!!!!!!!!!!*/
	}
	fmt.Println(d)

	// Example 2
	e := [2][2]int{{3, 5}, {7, 9}}
	fmt.Println(e)

	// Example 3
	var multiDimArr [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			multiDimArr[i][j] = i + j
		}
	}
	fmt.Println("2d: ", multiDimArr)
}

/*
 * Additional Resources:
 * https://dave.cheney.net/2014/10/04/that-trailing-comma

 */
