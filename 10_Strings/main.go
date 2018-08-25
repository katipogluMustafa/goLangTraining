package main

import "fmt"

func main(){
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo) )
	fmt.Printf("%d \n", rune(foo) )
}

/*
				Strings are collections of runes
				https://blog.golang.org/strings
*/
