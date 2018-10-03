package main

import "fmt"

func main() {
	var a, b , result uint32

	// Use This syntax to take input from different lines
	fmt.Scanf("%d\n%d",&a,&b)

	result = solveMeFirst(a,b)
	fmt.Println(result)
}

func solveMeFirst(a uint32, b uint32) uint32{
	return a + b
}