package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T - %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Println(int(myAge) + yourAge) // it needs to be converted to the same type
	fmt.Println(myAge + foo(yourAge)) // it needs to be converted to the same type
}
