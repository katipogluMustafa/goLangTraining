package main

import "fmt"

func main() {
	mySlice1 := make([]int, 0, 5)

	fmt.Println("------------")
	fmt.Println(mySlice1)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println("------------")

	for i := 0; i < 80; i++ {
		mySlice1 = append(mySlice1, i)
		fmt.Println("Len:", len(mySlice1), "Capacity:", cap(mySlice1), "Value: ", mySlice1[i])
	}
}

/*
 * make([]T, length	, capacity)
 * We use make to make a slice
 * make([]T, length	, capacity) is the same as new([100]int)[0:50]
 */
