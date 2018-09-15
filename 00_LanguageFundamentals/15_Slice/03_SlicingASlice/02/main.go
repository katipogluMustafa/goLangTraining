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

	// take slice of slice1 but capacity will be same as the length of it
	slice2 := slice1[2:4:4]
	inspectSlice(slice2)
	fmt.Println("*************************")

	// Now we still use the same underlying array
	// if we change anything, we are gonna see the change in both slices
	slice2[0] = "Changed"
	inspectSlice(slice2)
	inspectSlice(slice1)
	fmt.Println("*************************")

	// If we append anything to the slice2
	// since the capacity and length is the same
	// append() will create another underlying array
	// so we won't be using the same as slice1's underlying array
	slice2 = append(slice2, "Added")
	inspectSlice(slice1)
	inspectSlice(slice2)

	// Now if we change anything in slice2
	// we won't be seeing any change in slice1
	slice2[0] = "ChangedAgain"
	inspectSlice(slice2)
	inspectSlice(slice1)
	fmt.Println("*************************")
}

func inspectSlice(slice []string) {
	fmt.Printf("\nLength[%d] Capacity[%d]\n", len(slice), cap(slice))

	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
