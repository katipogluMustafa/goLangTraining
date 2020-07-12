package main

import "fmt"

func main(){

	// Escape Sequences
	fmt.Println("This is \n Escape Chacters \t lesson \t notes")
	fmt.Println("Even though you know these\\things")
	fmt.Println("Its funny to learn just like a beginner:)")

	// 'rune's implemented as unicode characters
	fmt.Println("Unicode of ğ is ",'ğ')
	fmt.Println("Unicode of a is ",'a')
	fmt.Println("Unicode of \\ is ", '\\')
	fmt.Println("Unicode of \\t is ", '\t')
}