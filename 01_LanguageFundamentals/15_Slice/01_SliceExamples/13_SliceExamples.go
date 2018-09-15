package main

import "fmt"

func main(){
	children := []int{1994,1996,1998}
	child := children[0:1]
	child[0] = 1990
	fmt.Println(children)
	fmt.Println(child)
}
/*
 * 	When we create a subslice we are getting back a reference to the original slice.
 * That means if we changed any value in the subslice the changes will be reflected to the original slice as well	.
 */