package main

import "fmt"

func main() {
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	inspectSlice(slice1)
	fmt.Println("*************************")

	// Parameters are [starting_index : (starting_index + length)]
	slice2 := slice1[2:4]
	inspectSlice(slice2)
	fmt.Println("*************************")

	slice2[0] = "Changed"
	inspectSlice(slice2)
	inspectSlice(slice1)
	fmt.Println("*************************")

	// Make a new slice big enough to hold elements of slice 2
	// copy the values over using the builtin copy function
	slice3 := make([]string, len(slice2))
	copy(slice3, slice2)
	inspectSlice(slice3)
	fmt.Println("*************************")
}

func inspectSlice(slice []string) {
	fmt.Printf("\nLength[%d] Capacity[%d]\n", len(slice), cap(slice))

	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
