package main

import "fmt"

func main() {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Iterate over the array
	// Displaying the value and address of each element.

	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%o]\n", v, &v, &friends[i])
	}

	fmt.Println(&friends[0])
	fmt.Println(&friends[1])
	fmt.Println(&friends[2])
	fmt.Println(&friends[3])
	fmt.Println(&friends[4])
}
