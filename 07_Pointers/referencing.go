package main

import "fmt"

func main(){
	a := 43

	fmt.Println("a - ",a)
	fmt.Println("&a -",&a)

	var b *int = &a
	fmt.Println("b",b)
	fmt.Println("*b",*b)
	fmt.Println("&b",&b)

	var c **int = &b
	fmt.Println("c",c)
	fmt.Println("*c",*c)
	fmt.Println("**c",**c)
	fmt.Println("&c",&c)

	fmt.Println(*&a)
	fmt.Println(*&b)
	fmt.Println(*&c)

}
