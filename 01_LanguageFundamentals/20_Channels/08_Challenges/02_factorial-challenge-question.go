package main

import "fmt"

func main(){
	f := factoriall(4)
	fmt.Println("Total:", f)
}

func factoriall(n int) int{
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
 # Challenge 1
	-- Use goroutines and channels to calculate factorial
 # Challenge 2
	-- Why might you want to use goroutines and channels to calculate factorial?
 */